package main

import (
	"context"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"fileshare/database/models"
)

func (inst *httpInstance) authProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Unable to complete user authentication.")
		inst.showErrorPage(w, r, http.StatusBadRequest, "Invalid session, please try again.")
		return
	}

	lFields := map[string]interface{}{
		"provider": provider,
		"email":    user.Email,
		"name":     user.Name,
		"id":       user.UserID,
	}

	account, err := inst.getAccountFromGothUser(user)
	if err != nil && !errors.Is(err, ErrAccountNotfound) {
		inst.logger.Error().Err(err).Fields(lFields).Msg("Unable to get oauth member from goth user.")
		inst.showErrorPage(w, r, http.StatusUnauthorized, "Unable to authenticate.")
		return
	}
	if err != nil && errors.Is(err, ErrAccountNotfound) {
		newAccount, errAcc := inst.createNewUserIfValidInviteToken(r, user)
		if errAcc != nil {
			if errors.Is(errAcc, ErrInvalidInviteToken) {
				inst.showErrorPage(w, r, http.StatusBadRequest, "Invalid invite token.")
				return
			}
			inst.showErrorPage(w, r, http.StatusInternalServerError, "Unable to create new account.")
			return
		}
		c := &http.Cookie{
			Name:   InviteTokenCookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, c)
		account = newAccount
	}

	inst.logger.Info().Fields(lFields).Str("account_id", account.ID).Msg("User authenticated.")
	sess, sessErr := inst.createAccountSession(account.ID)
	if sessErr != nil {
		inst.logger.Error().Err(sessErr).Fields(lFields).Msg("Unable to create session.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Unable to authenticate.")
		return
	}
	inst.createAccountSessionCookie(w, sess.ID)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (inst *httpInstance) authProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		lFields := map[string]interface{}{
			"provider": provider,
			"email":    user.Email,
			"name":     user.Name,
			"id":       user.UserID,
		}
		account, oErr := inst.getAccountFromGothUser(user)
		if oErr != nil {
			inst.logger.Error().Fields(lFields).Err(oErr).Msg("Unable to get oauth member from goth user.")
			inst.showErrorPage(w, r, http.StatusInternalServerError, "Unable to authenticate.")
			return
		}
		sess, sessErr := inst.createAccountSession(account.ID)
		if sessErr != nil {
			inst.logger.Error().Fields(lFields).Err(sessErr).Msg("Unable to create account session.")
			inst.showErrorPage(w, r, http.StatusInternalServerError, "Unable to authenticate.")
			return
		}
		inst.logger.Info().Fields(lFields).Str("account_id", account.ID).Msg("User authenticated.")
		inst.createAccountSessionCookie(w, sess.ID)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	inviteToken := r.URL.Query().Get(InviteTokenCookieName)

	if inviteToken != "" {
		inst.createInviteTokenCookie(w, inviteToken)
	}

	gothic.BeginAuthHandler(w, r)
}

func (inst *httpInstance) authLogout(w http.ResponseWriter, r *http.Request) {
	sessionID := inst.getSessionIDFromCookie(r)
	if sessionID != "" {
		_ = inst.deleteAccountSession(sessionID)
	}
	inst.deleteSessionIDCookie(w)
	_ = gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func getAccountFromContext(r *http.Request) *models.Account {
	user, ok := r.Context().Value(contextAccountKey).(*models.Account)
	if !ok {
		return nil
	}
	return user
}
