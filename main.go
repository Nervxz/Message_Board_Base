package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/nervxz/msg-board/internal/config"
	internal "github.com/nervxz/msg-board/internal/database"
	"github.com/nervxz/msg-board/internal/handlers"
	"github.com/nervxz/msg-board/internal/utils"
	"github.com/redis/go-redis/v9"
)

func main() {
	s, err := newServer()
	if err != nil {
		log.Fatalf("fail to create server: %v", err)
	}

	// run the server in another goroutine
	go s.run()

	// using main thread to catch kill/shutdown signal and cleanup
	if err = s.waitShutdown(); err != nil {
		os.Exit(1)
		return
	}
	
	
}

type server struct {
	done       chan struct{}
	httpServer *http.Server
	cfg        *config.ServerConfig
	db         *sql.DB
	redis      *redis.Client
}


// this function will run the server and listen to the port due to htpServer.ListenAndServe()
// then check the error whether it is nil or not due to golang error handling (ErrServerClosed)
func (s *server) run() {
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("fail to listen: %s\n", err)
	}
	close(s.done)
}


func (s *server) waitShutdown() error {
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down, stop accepting HTTP request...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Print("fail to shutdown HTTP server:", err)
		return err
	}

	// catching ctx.Done()
	select {
	case <-ctx.Done():
		log.Println("timeout during shutting down, continue to clean up")
	case <-s.done:
		log.Println("shut down in time")
		cancel()
	}

	if err := s.db.Close(); err != nil {
        log.Printf("fail to close DB connections: %v", err)
        return err
    }


	if err := s.redis.Close(); err != nil {
		log.Printf("fail to close Redis connections: %v", err)
		return err
	}

	log.Println("server shutdown successfully")
	return nil
}

func newServer() (*server, error) {
    cfg, err := config.Resolve()
    if err != nil {
        log.Fatalf("Failed to resolve configuration: %v", err)
    }

    db, err := connectDB(cfg.DB)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    
	

    err = internal.MigrateDB(db)
    if err != nil {
        log.Fatalf("Failed to migrate the database: %v", err)
    }

    redisClient, err := connectRedis(cfg.Redis)
    if err != nil {
        return nil, err
    }

    route := gin.Default()
    route.Use(gin.Logger())
    route.Use(gin.Recovery())

    route = handlers.Setup(route, db, redisClient)

    return &server{
        done:       make(chan struct{}),
        httpServer: newHTTP(route),
        cfg:        cfg,
        db:         db,
        redis:      redisClient,
    }, nil
}


func newHTTP(route *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    utils.LoadEnvOrDefault("SERVER_ADDR", ":8080"),
		Handler: route.Handler(),
	}
}

func connectDB(cfg config.DBConfig) (*sql.DB, error) {
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

func connectRedis(cfg config.RedisConfig) (*redis.Client, error) {
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

