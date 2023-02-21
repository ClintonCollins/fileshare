package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/steam"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fileshare/database/models"
)

var (
	ErrAccountNotfound    = errors.New("account not found")
	ErrInvalidInviteToken = errors.New("invalid invite token")
)

func (inst *httpInstance) setupGoth() {
	steamProvider := steam.New(inst.steamAPIKey, inst.publicURL+"/auth/steam/callback")
	discordProvider := discord.New(inst.discordClientKey, inst.discordClientSecret,
		inst.publicURL+"/auth/discord/callback",
		discord.ScopeEmail,
		discord.ScopeIdentify)
	goth.UseProviders(discordProvider)
	goth.UseProviders(steamProvider)

	gothStore := sessions.NewCookieStore([]byte("FileShare"))
	gothStore.MaxAge(86400 * 30)
	gothStore.Options.Path = "/"
	gothStore.Options.HttpOnly = true
	gothStore.Options.Secure = false
	gothic.Store = gothStore
}

func (inst *httpInstance) createNewAccountAndOauthMember(user goth.User) (*models.Account, error) {
	tx, errTx := inst.db.Begin()
	if errTx != nil {
		return nil, errTx
	}
	defer func() {
		_ = tx.Rollback()
	}()
	newAccount := &models.Account{
		Email:       user.Email,
		IsSuperuser: false,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	nErr := newAccount.Insert(context.TODO(), tx, boil.Infer())
	if nErr != nil {
		inst.logger.Error().Err(nErr).Str("email", user.Email).Msg("Unable to create new account.")
		return nil, nErr
	}
	newOauthUser := gothUserToOauthUser(user)
	newOauthUser.AccountID = newAccount.ID
	newOauthUser.R = newOauthUser.R.NewStruct()
	newOauthUser.R.Account = newAccount
	oErr := newOauthUser.Insert(context.TODO(), tx, boil.Infer())
	if oErr != nil {
		inst.logger.Error().Err(oErr).Str("user_id", newAccount.ID).Msg("Unable to update oauth member.")
		return nil, oErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		inst.logger.Error().Err(commitErr).Str("user_id", newAccount.ID).Msg("Unable to commit transaction.")
		return nil, commitErr
	}

	return newAccount, nil
}

func (inst *httpInstance) getAccountFromGothUser(user goth.User) (*models.Account, error) {
	oauthMember, err := models.AuthProviders(
		qm.Where("provider_user_id = ?", user.UserID),
		qm.And("id_token = ?", user.IDToken),
		qm.Load(models.AuthProviderRels.Account),
	).One(context.TODO(), inst.db)
	if err != nil {
		return nil, ErrAccountNotfound
	}
	if oauthMember.R.Account == nil {
		return nil, ErrAccountNotfound
	}
	return oauthMember.R.Account, nil
}

func gothUserToOauthUser(user goth.User) *models.AuthProvider {
	return &models.AuthProvider{
		Provider:          user.Provider,
		Email:             user.Email,
		Name:              user.Name,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		NickName:          user.NickName,
		Description:       user.Description,
		ProviderUserID:    user.UserID,
		AvatarURL:         user.AvatarURL,
		Location:          user.Location,
		AccessToken:       user.AccessToken,
		AccessTokenSecret: user.AccessTokenSecret,
		RefreshToken:      user.RefreshToken,
		ExpiresAt:         user.ExpiresAt,
		IDToken:           user.IDToken,
	}
}

func (inst *httpInstance) createAccountSession(accountID string) (*models.Session, error) {
	memberSession := &models.Session{
		AccountID: accountID,
		ExpiresAt: time.Now().AddDate(0, 0, 14),
	}
	err := memberSession.Insert(context.TODO(), inst.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return memberSession, nil
}

func (inst *httpInstance) createAccountSessionCookie(w http.ResponseWriter, sessionID string) {
	encodedCookie, eErr := inst.secureCookie.Encode(SessionCookieName, sessionID)
	if eErr != nil {
		inst.logger.Error().Err(eErr).Msg("")
		return
	}

	sessionCookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    encodedCookie,
		Path:     "/",
		Expires:  time.Now().AddDate(0, 0, 14),
		HttpOnly: true,
		Secure:   inst.useSSL || inst.useAutoSSL,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, sessionCookie)
}

func getSessionCookie(r *http.Request) (*http.Cookie, error) {
	sessionCookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return nil, err
	}
	return sessionCookie, nil
}

func (inst *httpInstance) updateSessionIfCloseToExpiration(session *models.Session) {
	if time.Until(session.ExpiresAt) < time.Hour*24*7 {
		session.ExpiresAt = time.Now().AddDate(0, 0, 14)
		_, err := session.Update(context.TODO(), inst.db, boil.Infer())
		if err != nil {
			inst.logger.Error().Str("session_id", session.ID).Err(err).Msg("Unable to update session expiration.")
		}
	}
}

func (inst *httpInstance) updateSessionCookieExpiration(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := getSessionCookie(r)
	if err != nil {
		return
	}
	newCookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionCookie.Value,
		Path:     "/",
		Expires:  time.Now().AddDate(0, 0, 14),
		HttpOnly: true,
		Secure:   inst.useSSL || inst.useAutoSSL,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, newCookie)
}

func (inst *httpInstance) getSessionIDFromCookie(r *http.Request) string {
	sessionCookie, err := getSessionCookie(r)
	if err != nil {
		return ""
	}
	var sessionID string
	decodeErr := inst.secureCookie.Decode(SessionCookieName, sessionCookie.Value, &sessionID)
	if decodeErr != nil {
		return ""
	}
	return sessionID
}

func (inst *httpInstance) deleteSessionIDCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:   SessionCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
}

func (inst *httpInstance) deleteAccountSession(sessionID string) error {
	_, err := models.Sessions(qm.Where("id = ?", sessionID)).DeleteAll(context.TODO(), inst.db)
	return err
}

func (inst *httpInstance) createInviteTokenCookie(w http.ResponseWriter, inviteToken string) {
	encodedCookie, eErr := inst.secureCookie.Encode(InviteTokenCookieName, inviteToken)
	if eErr != nil {
		inst.logger.Error().Err(eErr).Msg("Unable to encode secure cookie.")
		return
	}

	sessionCookie := &http.Cookie{
		Name:     InviteTokenCookieName,
		Value:    encodedCookie,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   inst.useSSL || inst.useAutoSSL,
	}

	http.SetCookie(w, sessionCookie)
}

func (inst *httpInstance) getInviteTokenFromCookie(r *http.Request) string {
	inviteCookie, err := r.Cookie(InviteTokenCookieName)
	if err != nil {
		inst.logger.Warn().Err(err).Msg("Unable to get invite token cookie.")
		return ""
	}
	var inviteToken string
	decodeErr := inst.secureCookie.Decode(InviteTokenCookieName, inviteCookie.Value, &inviteToken)
	if decodeErr != nil {
		inst.logger.Error().Err(decodeErr).Msg("Unable to decode secure cookie.")
		return ""
	}
	return inviteToken
}

func (inst *httpInstance) getInvitationFromToken(inviteToken string) (*models.Invitation, error) {
	invite, errInvite := models.Invitations(
		models.InvitationWhere.ID.EQ(inviteToken),
		models.InvitationWhere.Active.EQ(true),
		models.InvitationWhere.ExpiresAt.GT(time.Now()),
	).One(context.Background(), inst.db)
	if errInvite != nil {
		inst.logger.Error().Err(errInvite).Str("invite_token", inviteToken).Msg("Unable to get invitation from token.")
		return nil, errInvite
	}
	return invite, nil
}

func (inst *httpInstance) createNewUserIfValidInviteToken(r *http.Request, user goth.User) (*models.Account, error) {
	inviteToken := inst.getInviteTokenFromCookie(r)
	if inviteToken == "" {
		inst.logger.Warn().Msg("Unable to get cookie.")
		return nil, ErrInvalidInviteToken
	}
	invite, errInvite := inst.getInvitationFromToken(inviteToken)
	if errInvite != nil {
		return nil, ErrInvalidInviteToken
	}
	invite.Active = false
	_, errUpdate := invite.Update(context.Background(), inst.db, boil.Infer())
	if errUpdate != nil {
		inst.logger.Error().Err(errUpdate).Str("email", invite.Email).Str("id",
			invite.ID).Msg("Unable to deactivate invite.")
		return nil, errUpdate
	}

	newAccount, errCreate := inst.createNewAccountAndOauthMember(user)
	if errCreate != nil {
		return nil, errCreate
	}
	return newAccount, nil
}
