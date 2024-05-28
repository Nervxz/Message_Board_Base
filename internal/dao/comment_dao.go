package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/nervxz/msg-board/internal/dto"
)

func NewCommentDAO(db *sql.DB) *CommentDAO {
	return &CommentDAO{db: db}
}

type CommentDAO struct {
	db *sql.DB
}

// Takes a variable number of dto.Comment objects and returns a slice of created comments and an error
func (d CommentDAO) Create(comments ...dto.Comment) ([]dto.Comment, error) {
	if len(comments) == 0 {
		return nil, nil
	}
	// 	Generate the insert query for the number of comments provided
	query := d.genInsertSQL(len(comments))
	var args []any
	for _, c := range comments {
		args = append(args, c.By, c.ID, c.Content, time.Now())
	}
	// Insert the comments into the database
	res, err := d.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("fail to create comments: %w", err)
	}
	// Create a slice of int to store the created comment IDs
	ids := make([]int, 0, len(comments))
	for res.Next() {
		var id int
		if err = res.Scan(&id); err != nil {
			return nil, fmt.Errorf("fail to scan for created comment ID: %w", err)
		}
		// Append the created comment ID to the ids slice
		ids = append(ids, id)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning created comment IDs, err=%w", err)
	}

	for i, c := range comments {
		c.ID = ids[i]
	}

	return comments, nil
}

func (d CommentDAO) genInsertSQL(numComments int) string {
	var sb strings.Builder
	sb.WriteString(`insert into msg_board.comments(by, topic_id, body, created_at) values `)
	genSQLParams(numComments, &sb, 4)
	sb.WriteString(` returning id;`)
	return sb.String()
}
