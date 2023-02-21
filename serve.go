package main

import (
	"crypto/tls"
	"database/sql"
	"embed"
	"html/template"
	"io/fs"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/gorilla/securecookie"
	"github.com/rs/zerolog"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"

	"fileshare/pkg/configuration"
)

var (
	//go:embed templates/*
	templates embed.FS

	//go:embed frontend/dist/*
	staticFiles embed.FS
)

type httpInstance struct {
	devMode              bool
	db                   *sql.DB
	logger               zerolog.Logger
	mainTemplates        map[string]*template.Template
	secureCookie         *securecookie.SecureCookie
	fileStoragePath      string
	secureCookieHashKey  string
	secureCookieBlockKey string
	hostAddress          string
	hostPort             string
	publicURL            string
	discordClientKey     string
	discordClientSecret  string
	steamAPIKey          string
	useAutoSSL           bool
	useSSL               bool
	sslKeyFile           string
	sslCertFile          string
	shortIDGenerator     *shortid.Shortid
	staticFileHandler    http.Handler
	httpServer           *http.Server
}

const (
	// SessionCookieName is the name of the cookie used to store the session.
	SessionCookieName = "session_id"
	// InviteTokenCookieName is the name of the cookie used to store the invite token.
	InviteTokenCookieName = "invite_token"
)

func getHTTPServerInstance(config *configuration.Configuration, db *sql.DB, logger zerolog.Logger) (*httpInstance,
	error) {
	inst := &httpInstance{
		devMode:              config.DevMode,
		db:                   db,
		logger:               logger,
		mainTemplates:        make(map[string]*template.Template),
		fileStoragePath:      config.FileStoragePath,
		secureCookieHashKey:  config.SecureCookieHashKey,
		secureCookieBlockKey: config.SecureCookieBlockKey,
		hostAddress:          config.HostAddress,
		hostPort:             config.HostPort,
		publicURL:            config.PublicURL,
		discordClientKey:     config.DiscordClientKey,
		discordClientSecret:  config.DiscordClientSecret,
		steamAPIKey:          config.SteamAPIKey,
		useAutoSSL:           config.UseAutoSSL,
		useSSL:               config.UseSSL,
		sslKeyFile:           config.SSLKeyFile,
		sslCertFile:          config.SSLCertFile,
	}
	staticFileHandler, errStatic := getStaticFileHandler(config.DevMode)

	httpServer := &http.Server{
		Addr:                         net.JoinHostPort(inst.hostAddress, inst.hostPort),
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  time.Second * 60,
		ReadHeaderTimeout:            time.Second * 60,
		WriteTimeout:                 time.Second * 60,
		IdleTimeout:                  time.Second * 60,
	}
	inst.httpServer = httpServer

	if errStatic != nil {
		return nil, errStatic
	}
	inst.staticFileHandler = staticFileHandler
	if !config.DevMode {
		temps, err := inst.buildTemplatesMap()
		if err != nil {
			return nil, err
		}
		inst.mainTemplates = temps
	}
	if !strings.HasPrefix(inst.publicURL, "https://") && !strings.HasPrefix(inst.publicURL, "http://") {
		return nil, errors.New("Public URL must start with http:// or https://")
	}
	inst.secureCookie = securecookie.New([]byte(inst.secureCookieHashKey), []byte(inst.secureCookieBlockKey))
	inst.setupGoth()
	shortID, err := shortid.New(1, shortid.DefaultABC, 4)
	if err != nil {
		return nil, err
	}
	inst.shortIDGenerator = shortID
	return inst, nil
}

func getCurrentFileDirectory() string {
	var staticDirectoryPath string
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		staticDirectoryPath = "frontend/dist"
	} else {
		staticDirectoryPath = filepath.Dir(currentFile)
	}
	return staticDirectoryPath
}

func getStaticFileHandler(devMode bool) (http.Handler, error) {
	if !devMode {
		sfs, err := fs.Sub(staticFiles, "frontend/dist")
		if err != nil {
			return http.Handler(nil), err
		}
		return http.StripPrefix("/static/", http.FileServer(http.FS(sfs))), nil
	}
	return http.StripPrefix("/static/",
		http.FileServer(http.Dir(filepath.Join(getCurrentFileDirectory(), "frontend/dist")))), nil
}

func (inst *httpInstance) Listen() error {
	r := inst.buildRouter()

	inst.httpServer.Handler = r

	if inst.useAutoSSL {
		_ = os.MkdirAll(".fileshare_cache", os.ModePerm)
		autoTLSManager := &autocert.Manager{
			Prompt: autocert.AcceptTOS,
			Cache:  autocert.DirCache(".fileshare_cache"),
		}

		inst.httpServer.TLSConfig = &tls.Config{
			GetCertificate: autoTLSManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		}

		return inst.httpServer.ListenAndServeTLS("", "")
	}

	if inst.useSSL {
		return inst.httpServer.ListenAndServeTLS(inst.sslCertFile, inst.sslKeyFile)
	}

	return inst.httpServer.ListenAndServe()
}