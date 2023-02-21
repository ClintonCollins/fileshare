package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/volatiletech/null/v8"
)

type FileUploadResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	URL     string `json:"url"`
}

type FileTemplateObject struct {
	ID                string
	Name              string
	Slug              string
	ShortID           string
	Size              int64
	MimeType          string
	IsImage           bool
	IsVideo           bool
	IsAudio           bool
	URL               string
	PasswordProtected bool
	FileGroupID       null.String
	CreatedAt         time.Time
}

func (inst *httpInstance) getShortIDURLForFile(shortID string) string {
	return fmt.Sprintf("%s/f/%s", inst.publicURL, shortID)
}

func (inst *httpInstance) debugErrorPage(w http.ResponseWriter, r *http.Request) {
	inst.showErrorPage(w, r, http.StatusInternalServerError, "This is a debug error page.")
}

func (inst *httpInstance) showErrorPage(w http.ResponseWriter, r *http.Request, httpStatusCode int,
	errorMessage string) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"errorMessage":   errorMessage,
		"title":          "FileShare Error",
	}
	w.WriteHeader(httpStatusCode)
	inst.renderTemplate(w, data, "templates/error.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) get404(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare Page Not Found",
	}
	w.WriteHeader(http.StatusNotFound)
	inst.renderTemplate(w, data, "templates/404.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) getIndex(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare Upload",
	}
	inst.renderTemplate(w, data, "templates/index.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) getLogin(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	if currentAccount != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare Login",
	}
	inviteToken := r.URL.Query().Get(InviteTokenCookieName)

	if inviteToken != "" {
		inst.createInviteTokenCookie(w, inviteToken)
	}
	inst.renderTemplate(w, data, "templates/login.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) getAbout(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare About",
	}
	inst.renderTemplate(w, data, "templates/about.gohtml", "templates/layouts/main.gohtml")
}
