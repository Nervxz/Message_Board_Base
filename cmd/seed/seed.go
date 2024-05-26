package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nervxz/msg-board/internal/config"
	"github.com/nervxz/msg-board/internal/dao"
	"github.com/nervxz/msg-board/internal/dto"
	"github.com/nervxz/msg-board/internal/utils"
)

const numUsers = 10

func main() {
	db := connect()
	users, err := genUsers(db)
	if err != nil {
		log.Fatal(err)
		return
	}

	topics, err := genTopics(db, users)
	if err != nil {
		log.Fatal(err)
		return
	}

	genComments(db, users, topics)
	genVotes(db, users, topics)
}

func toJson(u dto.User) string {
	s, _ := json.Marshal(u)
	return string(s)
}

func genUsers(db *sql.DB) ([]dto.User, error) {
	d := dao.NewUserDAO(db)
	users := prepareUsers()
	newUsers, existedUsers, err := splitExistedUsers(d, users)
	if err != nil {
		return nil, err
	}

	createdUsers, err := d.Create(newUsers...)
	if err != nil {
		log.Printf("fail to create users, err=%v", err)
		return nil, nil
	}

	return append(existedUsers, createdUsers...), nil
}

func splitExistedUsers(d *dao.UserDAO, users []dto.User) ([]dto.User, []dto.User, error) {
	names := make([]string, 0, len(users))
	for _, u := range users {
		names = append(names, u.Username)
	}

	existedUsers, err := d.FindByUsername(names...)
	if err != nil {
		return nil, nil, err
	}

	seen := make(map[string]struct{})
	for _, u := range existedUsers {
		seen[u.Username] = struct{}{}
	}

	end := 0
	for _, u := range users {
		if _, existed := seen[u.Username]; !existed {
			users[end] = u
			end++
		}
	}
	users = users[:end]
	return users, existedUsers, nil
}

func prepareUsers() []dto.User {
	pass := utils.HashPass("password")
	users := make([]dto.User, 0, numUsers)
	for i := range numUsers {
		u := dto.User{
			Username: "user" + strconv.Itoa(i),
			Password: pass,
		}
		users = append(users, u)
	}
	return users
}

func genTopics(db *sql.DB, users []dto.User) ([]dto.Topic, error) {
	const minTopic = 2
	const maxTopic = 5
	topics := make([]dto.Topic, 0, len(users)*minTopic)
	now := time.Now()
	for _, u := range users {
		n := randInt(minTopic, maxTopic)
		for i := range n {
			topics = append(topics, dto.Topic{
				By:    u.ID,
				Title: fmt.Sprintf("topic %d by user %d at %d", i+1, u.ID, now.Unix()),
				Body:  "random string at " + now.String(),
			})
		}
	}

	d := dao.NewTopicDAO(db)
	created, err := d.Create(topics...)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func genVotes(db *sql.DB, users []dto.User, topics []dto.Topic) {

}

func genComments(db *sql.DB, users []dto.User, topics []dto.Topic) {
}

func connect() *sql.DB {
	url := utils.LoadEnvOrDefault(config.EnvDBURL, config.DefaultDBURL)
	db, err := config.ConnectDB(config.DBConfig{
		URL: url,
	})

	if err != nil {
		log.Fatalf("fail to connect to DB, url=%v, err=%v", url, err)
	}

	return db
}

func randInt(min, max int) int {
	return rand.Int()%(max-min) + min
}
