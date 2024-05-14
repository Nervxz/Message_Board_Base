package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type dependencies struct {
	db    *sql.DB
	redis *redis.Client
}

func Setup(g *gin.Engine, db *sql.DB, redis *redis.Client) *gin.Engine {
	deps := &dependencies{
		db:    db,
		redis: redis,
	}
	// setup routes
	setupTopics(g.Group("/topics"), deps)
	setupComments(g.Group("/comments"), deps)
	setupUsers(g.Group("/users"), deps)


	return g
}
