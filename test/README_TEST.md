# How To Test ?

- Test Query Database by open PostGreSQL and input those Query below to run

Test Relationship Between Users and Topics

```bash
SELECT Users.Username, Topics.Title
FROM Users
INNER JOIN Topics ON Users.UserID = Topics.UserID;
```

Test Relationship Between Users and Comments

```bash
SELECT Users.Username, Comments.Comment
FROM Users
INNER JOIN Comments ON Users.UserID = Comments.UserID;

```

Test Relationship Between Topics and Comments

```bash
SELECT Topics.Title, Comments.Comment
FROM Topics
INNER JOIN Comments ON Topics.TopicID = Comments.TopicID;

```

Test Relationship Between Users, Topics, and Comments ( change userId for optional test)

```bash
SELECT Users.Username, Topics.Title, Comments.Comment, Comments.Comments_time, Topics.DatePublished
FROM Users
INNER JOIN Topics ON Users.UserID = Topics.UserID
INNER JOIN Comments ON Topics.TopicID = Comments.TopicID
WHERE Users.UserID = 1;
```

Test Count of Topics per User

```bash
SELECT Users.Username, COUNT(Topics.TopicID) AS TopicCount
FROM Users
LEFT JOIN Topics ON Users.UserID = Topics.UserID
GROUP BY Users.Username;

```

Test Count of Comments per Topic

```bash
SELECT Topics.Title, COUNT(Comments.CommentID) AS CommentCount
FROM Topics
LEFT JOIN Comments ON Topics.TopicID = Comments.TopicID
GROUP BY Topics.Title;

```

Test Count of Comments per User

```bash
SELECT Users.Username, COUNT(Comments.CommentID) AS CommentCount
FROM Users
LEFT JOIN Comments ON Users.UserID = Comments.UserID
GROUP BY Users.Username;

```

Test Count of Topics and Comments per User

```bash
SELECT Users.Username, COUNT(DISTINCT Topics.TopicID) AS TopicCount, COUNT(Comments.CommentID) AS CommentCount
FROM Users
LEFT JOIN Topics ON Users.UserID = Topics.UserID
LEFT JOIN Comments ON Users.UserID = Comments.UserID
GROUP BY Users.Username;

```

To get the latest comment for each topic

```bash
SELECT Topics.Title, Comments.Comment, MAX(Comments.Comments_time) AS LatestCommentTime
FROM Topics
LEFT JOIN Comments ON Topics.TopicID = Comments.TopicID
GROUP BY Topics.Title, Comments.Comment;

```
