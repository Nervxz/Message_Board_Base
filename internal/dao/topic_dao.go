package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/nervxz/msg-board/internal/dto"
)

func NewTopicDAO(db *sql.DB) *TopicDAO {
	return &TopicDAO{db: db}
}

type TopicDAO struct {
	db *sql.DB
}

func (d TopicDAO) Create(objs ...dto.Topic) ([]dto.Topic, error) {

	if len(objs) == 0 {
		return nil, nil
	}

	query := d.genInsertSQL(len(objs))
	var args []any
	for _, o := range objs {
		// Make sure the arguments matches the orders generated in genInsertSQL!!!
		args = append(args, o.Title, o.Body, o.By, time.Now())
	}

	res, err := d.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("fail to create topics: %w", err)
	}

	ids := make([]int, 0, len(objs))
	for res.Next() {
		var id int
		if err = res.Scan(&id); err != nil {
			return nil, fmt.Errorf("fail to scan for created topic ID: %w", err)
		}
		ids = append(ids, id)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning created topic IDs, err=%w", err)
	}

	for i, o := range objs {
		o.ID = ids[i]
	}

	return objs, nil
}

func (d TopicDAO) genInsertSQL(numTopics int) string {
	var sb strings.Builder
	sb.WriteString(`insert into msg_board.topics(title, body, by, created_at) values `)
	genSQLParams(numTopics, &sb, 4)
	sb.WriteString(` returning id;`)
	return sb.String()
}
