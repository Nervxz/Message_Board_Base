package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func MigrateDB(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            UserID SERIAL PRIMARY KEY,
            Username VARCHAR(255) NOT NULL,
            Password BYTEA NOT NULL,
            Registered_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

        CREATE TABLE IF NOT EXISTS Topics (
            TopicID SERIAL PRIMARY KEY,
            Title VARCHAR(255) NOT NULL,
            Body TEXT NOT NULL,
            DatePublished TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            UserID INTEGER REFERENCES Users(UserID)
        );

        CREATE TABLE IF NOT EXISTS Comments (
            CommentID SERIAL PRIMARY KEY,
            Comment TEXT NOT NULL,
            TopicID INTEGER REFERENCES Topics(TopicID),
            UserID INTEGER REFERENCES Users(UserID),
            Comments_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)

    return err
}
