package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupTopics(g *gin.RouterGroup, deps *dependencies) {
    h := &TopicHandler{deps: deps}
    h.bind(g)
}

type TopicHandler struct {
    deps *dependencies
}

func (h *TopicHandler) bind(g *gin.RouterGroup) {
    g.GET("/", h.getAllTopics)
    g.GET("/:id", h.getTopicByID)
    g.POST("/", h.createTopic)
    g.PUT("/:id", h.updateTopic)
}

func (h *TopicHandler) getAllTopics(gtx *gin.Context) {
    rows, err := h.deps.db.Query("SELECT TopicID, Title, Body, DatePublished, UserID FROM Topics")
    if err != nil {
        gtx.String(http.StatusInternalServerError, "Failed to query topics: %v", err)
        return
    }
    defer rows.Close()

    var topics []map[string]interface{}
    for rows.Next() {
        var topic = make(map[string]interface{})
        var topicID, userID int
        var title, body, datePublished string
        if err := rows.Scan(&topicID, &title, &body, &datePublished, &userID); err != nil {
            gtx.String(http.StatusInternalServerError, "Failed to scan topic: %v", err)
            return
        }
        topic["TopicID"] = topicID
        topic["Title"] = title
        topic["Body"] = body
        topic["DatePublished"] = datePublished
        topic["UserID"] = userID
        topics = append(topics, topic)
    }

    gtx.JSON(http.StatusOK, topics)
}

func (h *TopicHandler) getTopicByID(gtx *gin.Context) {
    id := gtx.Param("id")
    row := h.deps.db.QueryRow("SELECT TopicID, Title, Body, DatePublished, UserID FROM Topics WHERE TopicID = $1", id)
    var topicID, userID int
    var title, body, datePublished string
    if err := row.Scan(&topicID, &title, &body, &datePublished, &userID); err != nil {
        if err == sql.ErrNoRows {
            gtx.String(http.StatusNotFound, "Topic not found")
            return
        }
        gtx.String(http.StatusInternalServerError, "Failed to query topic: %v", err)
        return
    }
    gtx.JSON(http.StatusOK, gin.H{"TopicID": topicID, "Title": title, "Body": body, "DatePublished": datePublished, "UserID": userID})
}

func (h *TopicHandler) createTopic(gtx *gin.Context) {
    var topic struct {
        Title   string `json:"Title"`
        Body    string `json:"Body"`
        UserID  int    `json:"UserID"`
    }
    if err := gtx.BindJSON(&topic); err != nil {
        gtx.String(http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Insert the topic into the database
    _, err := h.deps.db.Exec("INSERT INTO Topics (Title, Body, UserID) VALUES ($1, $2, $3)", topic.Title, topic.Body, topic.UserID)
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