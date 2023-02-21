package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fileshare/database/models"
)

func (inst *httpInstance) adminGetInvitations(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare Invitations",
	}
	queryMods := []qm.QueryMod{
		qm.OrderBy("invitation.created_at desc"),
		qm.GroupBy("invitation.id"),
	}

	queryM, queryMC, cParams := commonListQuery(r, 25, "email")
	queryMods = append(queryMods, queryM...)
	invitations, err := models.Invitations(queryMods...).All(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get invitations from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to get invitations from database.")
		return
	}

	count, err := models.Invitations(queryMC...).Count(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get invitations from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to get files from database.")
		return
	}

	params := url.Values{}
	search := cParams.Search
	if search != "" {
		params.Set("search", search)
	}
	params.Set("limit", strconv.Itoa(cParams.Limit))

	paginationData := genPaginationData(cParams, count, len(invitations), "/admin/invitations", params)

	flashMessage, flashErr := inst.getFlashMessage(w, r)
	if flashErr == nil {
		data["FlashMessage"] = flashMessage
	}
	data["invitations"] = invitations
	data["paginationData"] = paginationData
	data["search"] = search

	inst.renderTemplate(w, data, "templates/admin/invitations.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) adminPostCreateInvitation(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	errParseForm := r.ParseForm()
	if errParseForm != nil {
		inst.logger.Error().Err(errParseForm).Msg("Failed to parse form.")
		inst.showErrorPage(w, r, http.StatusBadRequest, "Invalid form data")
		return
	}

	formEmail := r.PostFormValue("email")
	if formEmail == "" {
		_ = inst.setFlashMessage(w, ErrorFlashAlert, "Email is required to create an invitation.")
		http.Redirect(w, r, "/admin/invitations", http.StatusSeeOther)
		return
	}

	newInvitation := &models.Invitation{
		Email:              formEmail,
		ExpiresAt:          time.Now().AddDate(0, 0, 7),
		Active:             true,
		CreatedByAccountID: null.StringFrom(currentAccount.ID),
	}

	errInsert := newInvitation.Insert(r.Context(), inst.db, boil.Infer())
	if errInsert != nil {
		inst.logger.Error().Err(errInsert).Msg("Failed to insert invitation into database.")
		_ = inst.setFlashMessage(w, ErrorFlashAlert, "Failed to create invitation.")
		http.Redirect(w, r, "/admin/invitations", http.StatusSeeOther)
		return
	}

	_ = inst.setFlashMessage(w, SuccessFlashAlert, "Invitation created successfully.")
	http.Redirect(w, r, "/admin/invitations", http.StatusSeeOther)
}

func (inst *httpInstance) adminPostDeleteInvitations(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	if !currentAccount.IsSuperuser {
		inst.showErrorPage(w, r, http.StatusForbidden, "Forbidden")
		return
	}
	err := r.ParseForm()
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to parse form.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete invitations.")
		return
	}
	invitationIDsFormValue, exist := r.PostForm["ids"]
	if !exist || len(invitationIDsFormValue) == 0 {
		http.Redirect(w, r, "/admin/invitations", http.StatusSeeOther)
		return
	}
	var invitationIDs []interface{}
	for _, id := range invitationIDsFormValue {
		invitationIDs = append(invitationIDs, id)
	}
	invitations, err := models.Invitations(
		qm.WhereIn("id IN ?", invitationIDs...),
	).All(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get invitations from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete invitations.")
		return
	}

	invitationsDeleted, errDeleteAll := invitations.DeleteAll(r.Context(), inst.db)
	if errDeleteAll != nil {
		inst.logger.Error().Err(errDeleteAll).Msg("Failed to delete invitations from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete invitations.")
		return
	}

	plural := "invitations"
	if invitationsDeleted == 1 {
		plural = "invitation"
	}
	_ = inst.setFlashMessage(w, SuccessFlashAlert,
		fmt.Sprintf("Successfully deleted %d %s!", invitationsDeleted, plural))

	http.Redirect(w, r, "/admin/invitations", http.StatusSeeOther)
}

func (inst *httpInstance) adminGetAccounts(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare Accounts",
	}

	queryMods := []qm.QueryMod{
		qm.OrderBy("account.created_at desc"),
		qm.GroupBy("account.id"),
	}

	queryM, queryMC, cParams := commonListQuery(r, 25, "email", "display_name")
	queryMods = append(queryMods, queryM...)
	accounts, err := models.Accounts(queryMods...).All(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get accounts from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to get accounts from database.")
		return
	}

	count, err := models.Accounts(queryMC...).Count(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get accounts from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to get accounts from database.")
		return
	}

	params := url.Values{}
	search := cParams.Search
	if search != "" {
		params.Set("search", search)
	}
	params.Set("limit", strconv.Itoa(cParams.Limit))

	paginationData := genPaginationData(cParams, count, len(accounts), "/admin/accounts", params)

	flashMessage, flashErr := inst.getFlashMessage(w, r)
	if flashErr == nil {
		data["FlashMessage"] = flashMessage
	}
	data["accounts"] = accounts
	data["paginationData"] = paginationData
	data["search"] = search
	inst.renderTemplate(w, data, "templates/admin/accounts.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) adminPostDeleteAccounts(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	if !currentAccount.IsSuperuser {
		inst.showErrorPage(w, r, http.StatusForbidden, "Forbidden")
		return
	}
	err := r.ParseForm()
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to parse form.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete accounts.")
		return
	}
	accountIDsFormValue, exist := r.PostForm["ids"]
	if !exist || len(accountIDsFormValue) == 0 {
		http.Redirect(w, r, "/admin/accounts", http.StatusSeeOther)
		return
	}
	var accountIDs []interface{}
	for _, id := range accountIDsFormValue {
		accountIDs = append(accountIDs, id)
	}
	accounts, err := models.Accounts(
		qm.WhereIn("id IN ?", accountIDs...),
		qm.And("is_superuser = ?", false),
		qm.And("id != ?", currentAccount.ID),
	).All(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get accounts from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete accounts.")
		return
	}

	accountsDeleted, errDeleteAll := accounts.DeleteAll(r.Context(), inst.db)
	if errDeleteAll != nil {
		inst.logger.Error().Err(errDeleteAll).Msg("Failed to delete accounts from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete accounts.")
		return
	}

	plural := "accounts"
	if accountsDeleted == 1 {
		plural = "account"
	}
	_ = inst.setFlashMessage(w, SuccessFlashAlert, fmt.Sprintf("Successfully deleted %d %s!", accountsDeleted, plural))

	http.Redirect(w, r, "/admin/accounts", http.StatusSeeOther)
}
