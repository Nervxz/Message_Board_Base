# How To Test ?

- Test Query Database by open PostGreSQL and input those Query below to run

- Insert Users

```bash
INSERT INTO Users (Username, Password, Registeredtime)
VALUES ('username', 'password', CURRENT_TIMESTAMP),
		('username2', 'password2', CURRENT_TIMESTAMP),
		('username3', 'password3', CURRENT_TIMESTAMP),
		('username4', 'password4', CURRENT_TIMESTAMP),
		('username5', 'password5', CURRENT_TIMESTAMP),
		('username6', 'password6', CURRENT_TIMESTAMP);
```

- Insert Topics

```bash
INSERT INTO topics (topicid, title, body, datepublished, userid)
VALUES
    (1, 'First Topic', 'Body of first topic', '2024-05-14 10:00:00', 1),
    (2, 'Second Topic', 'Body of second topic', '2024-05-14 11:30:00', 2),
    (3, 'Third Topic', 'Body of third topic', '2024-05-14 12:45:00', 3),
	(4, 'First Topic', 'Body of first topic', '2024-05-14 10:00:00', 4),
    (5, 'Second Topic', 'Body of second topic', '2024-05-14 11:30:00', 5),
    (6, 'Third Topic', 'Body of third topic', '2024-05-14 12:45:00', 6);
```

- Insert Comments

```bash
INSERT INTO comments (commentid, comment, topicid, userid, commentstime)
VALUES
    (1, 'Interesting perspective', 1, 1, '2023-05-13 15:10:00'),
    (2, 'I have a different opinion on this', 1, 3, '2023-05-13 16:45:00'),
    (3, 'Thanks for the informative post', 2, 4, '2023-05-13 18:30:00'),
	(4, 'Interesting perspectiveasdasdsa', 1, 2, '2023-05-13 15:10:00'),
    (5, 'I have a different opinion on thisasdasdadasdd', 1, 3, '2023-05-13 16:45:00'),
    (6, 'Thanks for the informative postasdasdad', 2, 4, '2023-05-13 18:30:00');
```

<!--........................................................................................................................................ -->

- Test Relationship Between Users and Topics

```bash
SELECT Users.Username, Topics.Title
FROM Users
INNER JOIN Topics ON Users.UserID = Topics.UserID;
```

- Test Relationship Between Users and Comments

```bash
SELECT Users.Username, Comments.Comment
FROM Users
INNER JOIN Comments ON Users.UserID = Comments.UserID;

```

- Test Relationship Between Topics and Comments

```bash
SELECT Topics.Title, Comments.Comment
FROM Topics
INNER JOIN Comments ON Topics.TopicID = Comments.TopicID;

```

- Test Relationship Between Users, Topics, and Comments ( change userId for optional test)

```bash
SELECT Users.Username, Topics.Title, Comments.Comment, Comments.Comments_time, Topics.DatePublished
FROM Users
INNER JOIN Topics ON Users.UserID = Topics.UserID
INNER JOIN Comments ON Topics.TopicID = Comments.TopicID
WHERE Users.UserID = 1;
```

- Test Count of Topics per User

```bash
SELECT Users.Username, COUNT(Topics.TopicID) AS TopicCount
FROM Users
LEFT JOIN Topics ON Users.UserID = Topics.UserID
GROUP BY Users.Username;

```

- Test Count of Comments per Topic

```bash
SELECT Topics.Title, COUNT(Comments.CommentID) AS CommentCount
FROM Topics
LEFT JOIN Comments ON Topics.TopicID = Comments.TopicID
GROUP BY Topics.Title;

```

- Test Count of Comments per User

```bash
SELECT Users.Username, COUNT(Comments.CommentID) AS CommentCount
FROM Users
LEFT JOIN Comments ON Users.UserID = Comments.UserID
GROUP BY Users.Username;

```

- Test Count of Topics and Comments per User

```bash
SELECT Users.Username, COUNT(DISTINCT Topics.TopicID) AS TopicCount, COUNT(Comments.CommentID) AS CommentCount
FROM Users
LEFT JOIN Topics ON Users.UserID = Topics.UserID
LEFT JOIN Comments ON Users.UserID = Comments.UserID
GROUP BY Users.Username;

```

- To get the latest comment for each topic

```bash
SELECT Topics.Title, Comments.Comment, MAX(Comments.Comments_time) AS LatestCommentTime
FROM Topics
LEFT JOIN Comments ON Topics.TopicID = Comments.TopicID
GROUP BY Topics.Title, Comments.Comment;

```
