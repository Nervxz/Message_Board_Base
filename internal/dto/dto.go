package dto

import "time"

type User struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
}

type Vote struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	By        int       `json:"by,omitempty"`
	TopicID   int       `json:"topic_id,omitempty"`
}

type Topic struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	By           int    `json:"by,omitempty"`
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	CommentCount int    `json:"comment_count,omitempty"`
	Upvotes      int    `json:"upvotes,omitempty"`
}

type Comment struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	By      int    `json:"by,omitempty"`
	Content string `json:"content,omitempty"`
	TopicID int    `json:"topic_id,omitempty"`
}
