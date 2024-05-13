package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupUsers(g *gin.RouterGroup, deps *dependencies) {
	h := &UserHandler{deps: deps}
	h.bind(g)
}

type UserHandler struct {
	deps *dependencies
}

func (h *UserHandler) bind(g *gin.RouterGroup) {
	g.GET("/", h.getAllUsers)
	g.GET("/:id", h.getUserByID)
	g.POST("/", h.createUser)
	g.PUT("/:id", h.updateUser)
	g.DELETE("/:id", h.deleteUser)
}

func (h *UserHandler) getAllUsers(gtx *gin.Context) {
	gtx.String(http.StatusOK, "Retrieve all users")
}

func (h *UserHandler) getUserByID(gtx *gin.Context) {
	id := gtx.Param("id")
	gtx.String(http.StatusOK, "Retrieve user with ID: %s", id)
}

func (h *UserHandler) createUser(gtx *gin.Context) {
	gtx.String(http.StatusOK, "User created")
}

func (h *UserHandler) updateUser(gtx *gin.Context) {
	id := gtx.Param("id")
	gtx.String(http.StatusOK, "Updated user with ID: %s", id)
}

func (h *UserHandler) deleteUser(gtx *gin.Context) {
	id := gtx.Param("id")
	gtx.String(http.StatusOK, "Deleted user with ID: %s", id)
}
