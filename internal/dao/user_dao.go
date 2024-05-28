package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/nervxz/msg-board/internal/dto"
)

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{db: db}
}

type UserDAO struct {
	db *sql.DB
}

func (d UserDAO) FindByUsername(names ...string) ([]dto.User, error) {
	// If no names provided, return nil
	if len(names) == 0 {
		return nil, nil
	}
	// Construct the query where username matches provided names
	query := "select id, username, password, created_at from msg_board.users u where username = any ($1)"
	//
	arg := `{` + strings.Join(names, ",") + `}`
	res, err := d.db.Query(query, arg)

	if err != nil {
		return nil, fmt.Errorf("fail to find user by names, err=%w", err)
	}
	// Create a slice of dto.User to store the result in users
	users := make([]dto.User, 0, 4)
	// Loop through the result set and scan the result into the users slice
	for res.Next() {
		var u dto.User
		// Scan the result into the user struct
		if err = res.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("fail to scan query result into model, err=%w", err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (d UserDAO) Create(users ...dto.User) ([]dto.User, error) {
	if len(users) == 0 {
		return nil, nil
	}
	// Generate the insert query for the number of users provided
	query := d.genInsertSQL(len(users))
	// Create a slice of any to store the arguments for the query
	var args []any
	// loop through the users and append the arguments to the args slice
	for _, u := range users {
		// Make sure the arguments matches the orders generated in genInsertSQL!!!
		args = append(args, u.Username, u.Password, time.Now())
	}
	// Insert the users into the database
	res, err := d.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("fail to create users: %w", err)
	}
	// Create a slice of int to store the created user IDs
	ids := make([]int, 0, len(users))
	// Loop through the result set and scan the result into the ids slice
	for res.Next() {
		var id int
		if err = res.Scan(&id); err != nil {
			return nil, fmt.Errorf("fail to scan for created user ID: %w", err)
		}
		ids = append(ids, id)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("error while scanning created user IDs, err=%w", err)
	}
	// Loop through the users and assign the created IDs to the users
	for i, u := range users {
		u.ID = ids[i]
	}

	return users, nil
}

func (d UserDAO) genInsertSQL(numUsers int) string {
	var sb strings.Builder

	sb.WriteString(`insert into msg_board.users(username, password, created_at) values `)
	// Generate the SQL for the number of users provided
	genSQLParams(numUsers, &sb, 3)
	sb.WriteString(` returning id;`)
	return sb.String()
}
