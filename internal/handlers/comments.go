package handlers

import (
	"database/sql"
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
 // getAll is a handler that returns all comments
func (h *CommentHandler) getAll(gtx *gin.Context) {
    rows, err := h.deps.db.Query("SELECT CommentID, Comment, TopicID, UserID, CommentsTime FROM Comments")
    if err != nil {
        gtx.String(http.StatusInternalServerError, "Failed to query comments: %v", err)
        return
    }
    defer rows.Close()

    // create a slice of maps to store the comments
    var comments []map[string]interface{}
    for rows.Next() {
        var comment = make(map[string]interface{})
        var commentID, topicID, userID int
        var commentText, commentsTime string
        if err := rows.Scan(&commentID, &commentText, &topicID, &userID, &commentsTime); err != nil {
            gtx.String(http.StatusInternalServerError, "Failed to scan comment: %v", err)
            return
        }
        comment["CommentID"] = commentID
        comment["Comment"] = commentText
        comment["TopicID"] = topicID
        comment["UserID"] = userID
        comment["CommentsTime"] = commentsTime
        comments = append(comments, comment)
    }

    gtx.JSON(http.StatusOK, comments)
}
// getOne is a handler that returns a single comment
func (h *CommentHandler) getOne(gtx *gin.Context, id string) {
    row := h.deps.db.QueryRow("SELECT CommentID, Comment, TopicID, UserID, CommentsTime FROM Comments WHERE CommentID = $1", id)
    var commentID, topicID, userID int
    var commentText, commentsTime string
    if err := row.Scan(&commentID, &commentText, &topicID, &userID, &commentsTime); err != nil {
        if err == sql.ErrNoRows {
            gtx.String(http.StatusNotFound, "Comment not found")
            return
        }
        gtx.String(http.StatusInternalServerError, "Failed to query comment: %v", err)
        return
    }
    gtx.JSON(http.StatusOK, gin.H{"CommentID": commentID, "Comment": commentText, "TopicID": topicID, "UserID": userID, "CommentsTime": commentsTime})
}
