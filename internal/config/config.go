package config

import (
	"log"

	moho "github.com/joho/godotenv"
	"github.com/nervxz/msg-board/internal/utils"
)

// DefaultDBURL as setup by the docker compose file, for local dev only
const DefaultDBURL = `postgres://postgres:011991@localhost:15432/postgres`
const EnvDBURL = "DB_URL"

// defaultRedisURL as setup by the docker compose file, for local dev only
const defaultRedisURL = `localhost:16379`

// ServerConfig is to define the server configuration, how many config in this sever for example DB and Redis
type ServerConfig struct {
	DB    DBConfig
	Redis RedisConfig
}

type DBConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

func Resolve() (*ServerConfig, error) {
	err := moho.Load()
	if err != nil {
		log.Printf("WARNING: no .env file or the file could not be read (%v)", err)
	}

	return &ServerConfig{
		DB: DBConfig{
			URL: utils.LoadEnvOrDefault("DB_URL", DefaultDBURL),
		},

		Redis: RedisConfig{
			URL: utils.LoadEnvOrDefault("REDIS_URL", defaultRedisURL),
		},
	}, nil
}
