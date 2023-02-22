package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"fileshare/database"
	"fileshare/pkg/configuration"
	"fileshare/pkg/filesharecli"
	"fileshare/pkg/sharelogger"
)

func main() {
	if len(os.Args) > 1 {
		filesharecli.Start()
		return
	}

	c, errConfig := configuration.Get()
	if errConfig != nil {
		log.Fatal().Err(errConfig).Msg("Failed to load configuration.")
	}
	logger := sharelogger.GetLogger(c.DevMode)

	d, errDB := database.GetDB(&database.Config{
		PostgresHost:     c.PostgresHost,
		PostgresPort:     c.PostgresPort,
		PostgresUser:     c.PostgresUser,
		PostgresPassword: c.PostgresPassword,
		PostgresDatabase: c.PostgresDatabase,
	})
	if errDB != nil {
		logger.Fatal().Err(errDB).Msg("Failed to connect to database.")
	}

	httpServerInstance, errHTTP := getHTTPServerInstance(c, d, logger)
	if errHTTP != nil {
		logger.Fatal().Err(errHTTP).Msg("Failed to create HTTP server instance.")
	}

	gracefulShutdown := false

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	ctx, ctxCancel := context.WithCancel(context.Background())

	idleConnsClosed := make(chan struct{})

	go func() {
		<-osSignals
		ctxCancel()
		gracefulShutdown = true
		logger.Info().Msg("Received interrupt signal, shutting down...")
		ctxShutdown, shutdownCtxCancel := context.WithTimeout(context.Background(), time.Second*30)
		if httpServerInstance.redirectHTTPServer != nil {
			go func() {
				errShutdown := httpServerInstance.redirectHTTPServer.Shutdown(ctxShutdown)
				if errShutdown != nil {
					logger.Fatal().Err(errShutdown).Msg("Failed to shutdown redirect HTTP server.")
				}
			}()
		}
		errShutdown := httpServerInstance.httpServer.Shutdown(ctxShutdown)
		if errShutdown != nil {
			logger.Fatal().Err(errShutdown).Msg("Failed to shutdown HTTP server.")
		}
		shutdownCtxCancel()
		// Don't wait for redirect server to shut down.
		close(idleConnsClosed)
	}()

	// Startup automated job handling.
	go automatedProcessLooper(ctx, d, logger)

	go func() {
		// Enable https redirect.
		if c.UseHTTPSRedirect {
			logger.Info().Msgf("Redirect webserver started and listening for connections on %s",
				net.JoinHostPort(c.HostAddress, "80"))
			errListen := httpServerInstance.ListenHTTPRedirect()
			if errListen != nil {
				if !gracefulShutdown {
					logger.Fatal().Err(errListen).Msg("Failed to start redirect HTTP server.")
				}
			}
		}
	}()

	logger.Info().Msgf("Webserver started and listening for connections on %s with public URL %s.",
		net.JoinHostPort(c.HostAddress, c.HostPort), c.PublicURL)
	errListen := httpServerInstance.Listen()
	if errListen != nil {
		if !gracefulShutdown {
			logger.Fatal().Err(errListen).Msg("Failed to start HTTP server.")
		}
	}
	<-idleConnsClosed
}
