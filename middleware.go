package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fileshare/database/models"
)

type contextKey string

const (
	contextAccountKey contextKey = "account"
)

var (
	ignoredRequestPathPrefixes = []string{"/static/", "/favicon.ico", "/f/"}
)

func (inst *httpInstance) isIgnoredPath(r *http.Request) bool {
	for _, prefix := range ignoredRequestPathPrefixes {
		if strings.HasPrefix(r.URL.Path, prefix) {
			return true
		}
	}
	return false
}

func (inst *httpInstance) setUserContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if inst.isIgnoredPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		sessionID := inst.getSessionIDFromCookie(r)
		if sessionID == "" {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextAccountKey, nil)))
			return
		}
		session, err := models.Sessions(
			qm.Where("id = ?", sessionID),
			qm.Load(models.SessionRels.Account),
		).One(context.TODO(), inst.db)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextAccountKey, nil)))
			return
		}
		inst.updateSessionCookieExpiration(w, r)
		inst.updateSessionIfCloseToExpiration(session)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextAccountKey, session.R.Account)))
	})
}

func (inst *httpInstance) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if inst.isIgnoredPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		t1 := time.Now()
		defer func() {
			logFields := map[string]interface{}{
				"remote_address":   r.RemoteAddr,
				"url":              r.URL.Path,
				"protocol_version": r.Proto,
				"method":           r.Method,
				"user_agent":       r.Header.Get("User-Agent"),
				"bytes_written":    ww.BytesWritten(),
				// "response_header":        ww.Header(),
				"bytes_received":         r.Header.Get("Content-Length"),
				"status":                 ww.Status(),
				"requested_completed_in": time.Since(t1).String(),
				"current_account":        r.Context().Value(contextAccountKey),
			}
			inst.logger.Info().Str("action", "web-request").
				Timestamp().
				Fields(logFields).Msg("")
		}()
		next.ServeHTTP(ww, r)
	})
}

func (inst *httpInstance) isAuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if inst.isIgnoredPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		if getAccountFromContext(r) == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (inst *httpInstance) isSuperUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if inst.isIgnoredPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		currentAccount := getAccountFromContext(r)
		if currentAccount == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if !currentAccount.IsSuperuser {
			inst.showErrorPage(w, r, http.StatusForbidden, "You do not have permission to access this page.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
