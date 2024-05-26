package config

import (
	"context"
	"database/sql"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectDB(cfg DBConfig) (*sql.DB, error) {
	// open Go client
	cli, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		log.Printf("fail to connect to database: %v", err)
		return nil, err
	}

	// ping check
	if err = cli.PingContext(context.Background()); err != nil {
		log.Printf("fail to ping database: %v", err)
		return nil, err
	}

	log.Printf("connected to DB")
	return cli, nil
}

func ConnectRedis(cfg RedisConfig) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{Addr: cfg.URL})

	// ping check
	ping := cli.Ping(context.Background())
	if err := ping.Err(); err != nil {
		log.Printf("fail to ping redis: %v", err)
		return nil, err
	}

	log.Printf("connected to redis")
	return cli, nil
}
