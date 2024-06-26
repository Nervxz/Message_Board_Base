package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nervxz/msg-board/internal/model"
)

func setupTopics(g *gin.RouterGroup, deps *dependencies) {
	h := &TopicHandler{deps: deps}
	g.GET("/", h.getAllTopics)                    // Allow unauthenticated access to get all topics
	g.GET("/:id", h.getTopicByID)                 // Allow unauthenticated access to get topic details
	g.GET("/:id/comments", h.getCommentsForTopic) // Allow unauthenticated access to get comments for a topic

	g.Use(AuthMiddleware(deps.redis))
	h.bind(g)
}

type TopicHandler struct {
	deps *dependencies
}

// bind is a method that binds the handlers to the router group
func (h *TopicHandler) bind(g *gin.RouterGroup) {
	g.POST("/", h.createTopic)
	g.PUT("/:id", h.updateTopic)
	g.POST("/:id/upvote", h.upvoteTopic)

}

// getAllTopics is a handler that returns all topics and upvotes
// internal/handlers/topics.go

func (h *TopicHandler) getAllTopics(gtx *gin.Context) {
	rows, err := h.deps.db.Query(`
        SELECT 
            t.TopicID, 
            t.Title, 
            t.Body, 
            t.DatePublished, 
            t.UserID, 
            t.Upvotes,
            (SELECT COUNT(*) FROM Comments c WHERE c.TopicID = t.TopicID) AS CommentCount
        FROM Topics t
        ORDER BY t.Upvotes DESC, t.DatePublished DESC
    `)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to query topics: %v", err)
		return
	}
	defer rows.Close()

	var topics []model.Topic
	for rows.Next() {
		var c model.Topic
		if err := rows.Scan(&c.TopicID, &c.Title, &c.Body, &c.DatePublished, &c.UserID, &c.Upvotes, &c.CommentCount); err != nil {
			gtx.String(http.StatusInternalServerError, "Failed to scan topic: %v", err)
			return
		}

		topics = append(topics, c)
	}

	gtx.JSON(http.StatusOK, topics)
}

// getTopicByID is a handler that returns a single topic
func (h *TopicHandler) getTopicByID(gtx *gin.Context) {
	id := gtx.Param("id")
	row := h.deps.db.QueryRow("SELECT TopicID, Title, Body, DatePublished, UserID FROM Topics WHERE TopicID = $1", id)
	var c model.Topic
	if err := row.Scan(&c.TopicID, &c.Title, &c.Body, &c.DatePublished, &c.UserID); err != nil {
		if err == sql.ErrNoRows {
			gtx.String(http.StatusNotFound, "Topic not found")
			return
		}
		gtx.String(http.StatusInternalServerError, "Failed to query topic: %v", err)
		return
	}
	gtx.JSON(http.StatusOK, c)
}

// getCommentsForTopic is a handler that returns all comments for a specific topic
func (h *TopicHandler) getCommentsForTopic(gtx *gin.Context) {
	topicID := gtx.Param("id")
	rows, err := h.deps.db.Query("SELECT CommentID, Comment, TopicID, UserID, CommentsTime FROM Comments WHERE TopicID = $1", topicID)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to query comments: %v", err)
		return
	}
	defer rows.Close()

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

// createTopic is a handler that creates a new topic
func (h *TopicHandler) createTopic(gtx *gin.Context) {
	var topic struct {
		Title string `json:"Title"`
		Body  string `json:"Body"`
	}
	if err := gtx.BindJSON(&topic); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
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

	// Insert the topic into the database
	_, err = h.deps.db.Exec("INSERT INTO Topics (Title, Body, UserID) VALUES ($1, $2, $3)", topic.Title, topic.Body, userIDInt)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to create topic: %v", err)
		return
	}

	gtx.String(http.StatusOK, "Topic created successfully")
}

func (h *TopicHandler) updateTopic(gtx *gin.Context) {
	id := gtx.Param("id")
	var topic struct {
		Title string `json:"Title"`
		Body  string `json:"Body"`
	}
	if err := gtx.BindJSON(&topic); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update the topic in the database
	_, err := h.deps.db.Exec("UPDATE Topics SET Title = $1, Body = $2 WHERE TopicID = $3", topic.Title, topic.Body, id)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to update topic: %v", err)
		return
	}

	gtx.String(http.StatusOK, "Topic updated successfully")
}

// upvoteTopic is a handler that increments the upvote count for a topic
func (h *TopicHandler) upvoteTopic(gtx *gin.Context) {
	topicID := gtx.Param("id")

	// Get userID from context
	userID, exists := gtx.Get("userID")
	if !exists {
		gtx.String(http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Check if the user has already upvoted this topic
	var count int
	err := h.deps.db.QueryRow("SELECT COUNT(*) FROM Upvotes WHERE TopicID = $1 AND UserID = $2", topicID, userID).Scan(&count)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to query upvotes: %v", err)
		return
	}

	if count > 0 {
		gtx.String(http.StatusForbidden, "User has already upvoted this topic")
		return
	}

	// Insert the upvote into the database
	_, err = h.deps.db.Exec("INSERT INTO Upvotes (TopicID, UserID) VALUES ($1, $2)", topicID, userID)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to insert upvote: %v", err)
		return
	}

	// Increment the upvote count in the Topics table
	_, err = h.deps.db.Exec("UPDATE Topics SET Upvotes = Upvotes + 1 WHERE TopicID = $1", topicID)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to update upvote count: %v", err)
		return
	}

	gtx.String(http.StatusOK, "Upvote successful")
}
