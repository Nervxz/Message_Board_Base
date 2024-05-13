package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupComments(g *gin.RouterGroup, deps *dependencies) {
	h := &CommentHandler{deps: deps}
	h.bind(g)
}

type CommentHandler struct {
	deps *dependencies
}

func (h *CommentHandler) bind(g *gin.RouterGroup) {
	g.GET("/", h.getAll)
	g.GET("/:id", func(gtx *gin.Context) {
		id := gtx.Param("id")
		h.getOne(gtx, id) // TODO: parse ID into correct type
	})
}

func (h *CommentHandler) getAll(gtx *gin.Context) {
	gtx.String(http.StatusOK, "get all comments")
}

func (h *CommentHandler) getOne(gtx *gin.Context, id string) {
	gtx.String(http.StatusOK, "get detailed comment with id: %v", id)
}
