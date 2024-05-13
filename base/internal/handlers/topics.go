package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupTopics(g *gin.RouterGroup, deps *dependencies) {
	g.GET("/", func(gtx *gin.Context) {
		gtx.String(http.StatusOK, "get all topics")
	})

	g.GET("/:id", func(gtx *gin.Context) {
		id := gtx.Param("id")
		gtx.String(http.StatusOK, "get topic with id: %s", id)
	})

	g.POST("/", func(gtx *gin.Context) {
		// TODO: check auth sessions, validate data, ...
		gtx.String(http.StatusOK, "create a new topic")
	})

	g.PUT("/:id", func(gtx *gin.Context) {
		// TODO: check auth sessions, validate data, ...
		id := gtx.Param("id")
		gtx.String(http.StatusOK, "update topic with id: %s", id)
	})
}
