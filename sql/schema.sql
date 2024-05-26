/*
DB migration on service startup:
--------------------------------
- Django, Ruby on Rails, Spring Boot: only recommended for development.
- Hard to write a good migration script.
- In case of multi instances/distributed system, on new deployment, which server
instance will handle the migration?
*/

CREATE TABLE IF NOT EXISTS Users
(
    UserID         SERIAL PRIMARY KEY,
    Username       VARCHAR(255) NOT NULL,
    Password       BYTEA        NOT NULL,
    RegisteredTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Topics
(
    TopicID       SERIAL PRIMARY KEY,
    Title         VARCHAR(255) NOT NULL,
    Body          TEXT         NOT NULL,
    DatePublished TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UserID        INTEGER REFERENCES Users (UserID)
);

CREATE TABLE IF NOT EXISTS Comments
(
    CommentID    SERIAL PRIMARY KEY,
    Comment      TEXT NOT NULL,
    TopicID      INTEGER REFERENCES Topics (TopicID),
    UserID       INTEGER REFERENCES Users (UserID),
    CommentsTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Upvotes
(
    UpvoteID SERIAL PRIMARY KEY,
    TopicID  INTEGER REFERENCES Topics (TopicID),
    UserID   INTEGER REFERENCES Users (UserID),
    UNIQUE (TopicID, UserID)
);
