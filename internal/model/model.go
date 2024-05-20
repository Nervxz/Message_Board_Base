package model

import "time"

type Comment struct {
	TopicID      int       `json:"TopicID"`
	UserID       int       `json:"UserID"`
	CommentID    int       `json:"CommentID"`
	Comment      string    `json:"Comment"`
	CommentsTime time.Time `json:"CommentsTime"`
}

type Topic struct {
	TopicID       int       `json:"TopicID"`
	UserID        int       `json:"UserID"`
	Title         string    `json:"Title"`
	Body          string    `json:"Body"`
	DatePublished time.Time `json:"DatePublished"`
	Upvotes       int       `json:"Upvotes"`
}

type User struct {
	UserID         int       `json:"UserID"`
	Username       string    `json:"Username"`
	Password       string    `json:"Password"`
	RegisteredTime time.Time `json:"RegisteredTime"`
}
