package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*
var embeddedMigrations embed.FS

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func runMigrations(db *sql.DB) error {
	migrationFS := migrate.EmbedFileSystemMigrationSource{
		FileSystem: embeddedMigrations,
		Root:       "migrations",
	}
	totalMigrations, err := migrate.Exec(db, "postgres", migrationFS, migrate.Up)
	if totalMigrations > 0 {
		log.Info().Msgf("Applied %d migrations", totalMigrations)
	}
	return err
}

func GetDB(config *Config) (*sql.DB, error) {
	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDatabase,
	)
	db, errDB := sql.Open("postgres", postgresInfo)
	if errDB != nil {
		return nil, errDB
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}
	errMigrations := runMigrations(db)
	if errMigrations != nil {
		return nil, errMigrations
	}
	return db, nil
}
