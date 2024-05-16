package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nervxz/msg-board/internal/model"
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
		g.POST("/", h.createComment)
	})
}

// getAll is a handler that returns all comments
func (h *CommentHandler) getAll(gtx *gin.Context) {
	rows, err := h.deps.db.Query("SELECT CommentID, Comment, TopicID, UserID, CommentsTime FROM Comments")
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to query comments: %v", err)
		return
	}
	defer rows.Close()

	// create a slice of maps to store the comments
	var comments []model.Comment

	for rows.Next() {
		var c model.Comment
		if err := rows.Scan(&c.CommentID, &c.Comment, &c.TopicID, &c.UserID, &c.CommentsTime); err != nil {
			gtx.String(http.StatusInternalServerError, "Failed to scan comment: %v", err)
			return
		}
		comments = append(comments, c)
	}

	gtx.JSON(http.StatusOK, comments)
}

// getOne is a handler that returns a single comment
func (h *CommentHandler) getOne(gtx *gin.Context, id string) {
	row := h.deps.db.QueryRow("SELECT CommentID, Comment, TopicID, UserID, CommentsTime FROM Comments WHERE CommentID = $1", id)
	var c model.Comment

	if err := row.Scan(&c.CommentID, &c.Comment, &c.TopicID, &c.UserID, &c.CommentsTime); err != nil {
		if err == sql.ErrNoRows {
			gtx.String(http.StatusNotFound, "Comment not found")
			return
		}
		gtx.String(http.StatusInternalServerError, "Failed to query comment: %v", err)
		return
	}
	gtx.JSON(http.StatusOK, c)
}

// createComment is a handler that creates a new comment
func (h *CommentHandler) createComment(gtx *gin.Context) {
	var c model.Comment
	if err := gtx.BindJSON(c); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Insert the comment into the database
	_, err := h.deps.db.Exec("INSERT INTO Comments (Comment, TopicID, UserID) VALUES ($1, $2, $3)", c.Comment, c.TopicID, c.UserID)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to create comment: %v", err)
		return
	}

	gtx.String(http.StatusOK, "Comment created successfully")
}
