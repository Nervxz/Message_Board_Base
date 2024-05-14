package config

import (
	moho "github.com/joho/godotenv"
	"github.com/nervxz/msg-board/internal/utils"
)

// defaultDBURL as setup by the docker compose file, for local dev only
const defaultDBURL = `postgres://postgres:011991@localhost:5432/postgres`

// defaultRedisURL as setup by the docker compose file, for local dev only
const defaultRedisURL = `localhost:6379`


// ServerConfig is to define the server configuration, how many config in this sever for example DB and Redis
type ServerConfig struct {
	DB    DBConfig
	Redis RedisConfig
}
// Config the DBConfig which type of DB
type DBConfig struct {
	URL		 string
	Host     string
    Port     string
    User     string
    Password string
    DBName   string
}
// Config the RedisConfig which type of DB
type RedisConfig struct {
	URL string
	
}

// 
func Resolve() (*ServerConfig, error) {
	err := moho.Load()
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		DB: DBConfig{
			URL: 	  utils.LoadEnvOrDefault("DB_URL", defaultDBURL),
			Host:     utils.LoadEnvOrDefault("DB_HOST", "localhost"),
            Port:     utils.LoadEnvOrDefault("DB_PORT", "15432"),
            User:     utils.LoadEnvOrDefault("POSTGRES_USER", "postgres"),
            Password: utils.LoadEnvOrDefault("POSTGRES_PASSWORD", "011991"),
            DBName:   utils.LoadEnvOrDefault("DB_NAME", "postgres"),
		},

		
		Redis: RedisConfig{
			URL: utils.LoadEnvOrDefault("REDIS_URL", defaultRedisURL),
		},
	}, nil
}
