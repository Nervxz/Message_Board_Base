package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nervxz/msg-board/internal/dto"
)

func setupComments(g *gin.RouterGroup, deps *dependencies) {
	h := &CommentHandler{deps: deps}
	g.Use(AuthMiddleware(deps.redis)) // Protect the comment creation endpoint
	h.bind(g)
}

func (h *CommentHandler) bind(g *gin.RouterGroup) {
	g.POST("/", h.createComment)
	g.GET("/:topicID", h.getCommentsByTopicID) // Fetch comments for a specific topic
}

type CommentHandler struct {
	deps *dependencies
}

func (h *CommentHandler) createComment(gtx *gin.Context) {
	var c dto.Comment
	if err := gtx.BindJSON(&c); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		log.Printf("Error binding JSON: %v", err)
		return
	}

	log.Printf("Received comment: %+v", c)

	// Check if the topicId exists
	var exists bool
	err := h.deps.db.QueryRow("SELECT EXISTS(SELECT 1 FROM Topics WHERE TopicID = $1)", c.TopicID).Scan(&exists)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to check topicId: %v", err)
		return
	}
	if !exists {
		gtx.String(http.StatusBadRequest, "Invalid topicId")
		return
	}

	// Get userID from context
	userID, exists := gtx.Get("userID")
	if !exists {
		gtx.String(http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Convert userID to int
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to parse userID")
		return
	}

	// Insert the comment into the database
	_, err = h.deps.db.Exec("INSERT INTO Comments (Comment, TopicID, UserID) VALUES ($1, $2, $3)", c.Content, c.TopicID, userIDInt)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to create comment: %v", err)
		return
	}

	gtx.String(http.StatusOK, "Comment created successfully")
}

// getCommentsByTopicID is a handler that returns all comments for a specific topic
func (h *CommentHandler) getCommentsByTopicID(gtx *gin.Context) {
	topicID := gtx.Param("topicID")
	rows, err := h.deps.db.Query("SELECT CommentID, Comment, TopicID, UserID, CommentsTime FROM Comments WHERE TopicID = $1", topicID)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to query comments: %v", err)
		return
	}
	defer rows.Close()

	var comments []dto.Comment
	for rows.Next() {
		var c dto.Comment
		if err := rows.Scan(&c.ID, &c.Content, &c.TopicID, &c.By, &c.CreatedAt); err != nil {
			gtx.String(http.StatusInternalServerError, "Failed to scan comment: %v", err)
			return
		}
		comments = append(comments, c)
	}

	gtx.JSON(http.StatusOK, comments)
}
