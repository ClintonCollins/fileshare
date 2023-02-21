package main

import (
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
)

func (inst *httpInstance) buildRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(inst.setUserContextMiddleware)
	router.Use(inst.loggerMiddleware)

	// Auth not required
	router.Get("/login", inst.getLogin)
	router.Get("/f/{shortID}", inst.getFileDownload)

	// Auth Providers
	router.Get("/auth/{provider}/callback", inst.authProviderCallback)
	router.Get("/auth/{provider}", inst.authProvider)

	// Authenticated routes
	router.Group(func(r chi.Router) {
		r.Use(inst.isAuthenticatedMiddleware)

		r.Get("/", inst.getIndex)
		r.Get("/favicon*", func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = path.Join("/static/images/favicon", r.URL.Path)
			inst.staticFileHandler.ServeHTTP(w, r)
		})
		r.Get("/files", inst.getFileList)
		r.Get("/about", inst.getAbout)
		r.Post("/files/delete", inst.postHandleFileDelete)
		r.Post("/upload", inst.postHandleFileUpload)

		// Logout
		r.Get("/auth/logout", inst.authLogout)
	})

	// Superuser routes
	router.Group(func(r chi.Router) {
		r.Use(inst.isSuperUserMiddleware)

		r.Get("/admin/invitations", inst.adminGetInvitations)
		r.Post("/admin/invitations", inst.adminPostCreateInvitation)
		r.Post("/admin/invitations/delete", inst.adminPostDeleteInvitations)

		r.Get("/admin/accounts", inst.adminGetAccounts)
		r.Post("/admin/accounts/delete", inst.adminPostDeleteAccounts)
	})

	router.Get("/*", inst.get404)

	// Static files
	router.Handle("/static/*", inst.staticFileHandler)
	return router
}
