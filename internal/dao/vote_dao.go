package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/nervxz/msg-board/internal/dto"
)

func NewVoteDAO(db *sql.DB) *VoteDAO {
	return &VoteDAO{db: db}
}

type VoteDAO struct {
	db *sql.DB
}

func (d VoteDAO) Create(votes ...dto.Vote) ([]dto.Vote, error) {
	if len(votes) == 0 {
		return nil, nil
	}

	query := d.genInsertSQL(len(votes))
	var args []any
	for _, v := range votes {
		args = append(args, v.By, v.TopicID, time.Now())
	}

	res, err := d.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("fail to create votes: %w", err)
	}

	ids := make([]int, 0, len(votes))
	for res.Next() {
		var id int
		if err = res.Scan(&id); err != nil {
			return nil, fmt.Errorf("fail to scan for created vote ID: %w", err)
		}
		ids = append(ids, id)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning created vote IDs, err=%w", err)
	}

	for i, v := range votes {
		v.ID = ids[i]
	}

	return votes, nil
}

func (d VoteDAO) genInsertSQL(numVotes int) string {
	var sb strings.Builder
	sb.WriteString(`insert into msg_board.votes(by, topic_id, created_at) values `)
	genSQLParams(numVotes, &sb, 3)
	sb.WriteString(` returning id;`)
	return sb.String()
}
