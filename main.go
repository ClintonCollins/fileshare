package main

import (
	"context"
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

	go func() {
		<-osSignals
		ctxCancel()
		gracefulShutdown = true
		logger.Info().Msg("Received interrupt signal, shutting down...")
		ctxShutdown, _ := context.WithTimeout(context.Background(), time.Second*30)
		errShutdown := httpServerInstance.httpServer.Shutdown(ctxShutdown)
		if errShutdown != nil {
			logger.Fatal().Err(errShutdown).Msg("Failed to shutdown HTTP server.")
		}
	}()

	// Startup automated job handling.
	go automatedProcessLooper(ctx, d, logger)

	logger.Info().Msg("Webserver started and listening for connections.")
	errListen := httpServerInstance.Listen()
	if errListen != nil {
		if !gracefulShutdown {
			logger.Fatal().Err(errListen).Msg("Failed to start HTTP server.")
		}
	}
}
