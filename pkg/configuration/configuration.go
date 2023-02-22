package configuration

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Configuration struct {
	DevMode              bool   `env:"DEV_MODE"`
	PostgresHost         string `env:"POSTGRES_HOST,required"`
	PostgresPort         string `env:"POSTGRES_PORT,required"`
	PostgresUser         string `env:"POSTGRES_USER,required"`
	PostgresPassword     string `env:"POSTGRES_PASSWORD,required"`
	PostgresDatabase     string `env:"POSTGRES_DATABASE,required"`
	FileStoragePath      string `env:"FILE_STORAGE_PATH,required"`
	SecureCookieHashKey  string `env:"SECURE_COOKIE_HASH_KEY,required"`
	SecureCookieBlockKey string `env:"SECURE_COOKIE_BLOCK_KEY,required"`
	HostAddress          string `env:"HOST_ADDRESS" envDefault:""`
	HostPort             string `env:"HOST_PORT" envDefault:"8000"`
	PublicURL            string `env:"PUBLIC_URL,required"`
	DiscordClientKey     string `env:"DISCORD_CLIENT_KEY,required"`
	DiscordClientSecret  string `env:"DISCORD_CLIENT_SECRET,required"`
	SteamAPIKey          string `env:"STEAM_API_KEY,required"`
	UseAutoSSL           bool   `env:"USE_AUTO_SSL" envDefault:"false"`
	UseSSL               bool   `env:"USE_SSL" envDefault:"false"`
	SSLKeyFile           string `env:"SSL_KEY_FILE"`
	SSLCertFile          string `env:"SSL_CERT_FILE"`
	UseHTTPSRedirect     bool   `env:"USE_HTTPS_REDIRECT" envDefault:"false"`
}

func Get() (*Configuration, error) {
	config, errConfig := loadEnv()
	if errConfig != nil {
		return nil, errConfig
	}
	return config, nil
}

func loadEnv() (*Configuration, error) {
	config := &Configuration{}
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Debug().Err(errEnv).Msg("Failed to load .env file.")
	}
	err := env.Parse(config, env.Options{
		Prefix: "FILESHARE_",
	})
	if err != nil {
		return nil, err
	}
	return config, nil
}
