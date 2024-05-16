package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupUsers(g *gin.RouterGroup, deps *dependencies) {
	h := &UserHandler{deps: deps}
	h.bind(g)
}

type UserHandler struct {
	deps *dependencies
}

func (h *UserHandler) bind(g *gin.RouterGroup) {
	g.GET("/", h.getAllUsers)
	g.GET("/:id", h.getUserByID)
	g.POST("/", h.createUser)
	g.PUT("/:id", h.updateUser)
	g.DELETE("/:id", h.deleteUser)
}

func (h *UserHandler) getAllUsers(gtx *gin.Context) {
	rows, err := h.deps.db.Query("SELECT UserID, Username, RegisteredTime FROM Users")
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to query users: %v", err)
		return
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var user = make(map[string]interface{})
		var userID int
		var username string
		var registeredTime string

		if err := rows.Scan(&userID, &username, &registeredTime); err != nil {
			gtx.String(http.StatusInternalServerError, "Failed to scan user: %v", err)
			return
		}
		user["UserID"] = userID
		user["Username"] = username
		user["RegisteredTime"] = registeredTime

		users = append(users, user)
	}

	gtx.JSON(http.StatusOK, users)
}

func (h *UserHandler) getUserByID(gtx *gin.Context) {
	id := gtx.Param("id")
	row := h.deps.db.QueryRow("SELECT UserID, Username, RegisteredTime FROM Users WHERE UserID = $1", id)
	var userID int
	var username, registeredTime string
	if err := row.Scan(&userID, &username, &registeredTime); err != nil {
		if err == sql.ErrNoRows {
			gtx.String(http.StatusNotFound, "User not found")
			return
		}
		gtx.String(http.StatusInternalServerError, "Failed to query user: %v", err)
		return
	}
	gtx.JSON(http.StatusOK, gin.H{"UserID": userID, "Username": username, "RegisteredTime": registeredTime})
}

func (h *UserHandler) createUser(gtx *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := gtx.BindJSON(&user); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Insert the user into the database
	_, err := h.deps.db.Exec("INSERT INTO Users (Username, Password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to create user: %v", err)
		return
	}

	gtx.String(http.StatusOK, "User created successfully")
}

func (h *UserHandler) updateUser(gtx *gin.Context) {
	id := gtx.Param("id")
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := gtx.BindJSON(&user); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update the user in the database
	_, err := h.deps.db.Exec("UPDATE Users SET Username = $1, Password = $2 WHERE UserID = $3", user.Username, user.Password, id)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to update user: %v", err)
		return
	}

	gtx.String(http.StatusOK, "User updated successfully")
}

func (h *UserHandler) deleteUser(gtx *gin.Context) {
	id := gtx.Param("id")

	// Delete the user from the database
	_, err := h.deps.db.Exec("DELETE FROM Users WHERE UserID = $1", id)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to delete user: %v", err)
		return
	}

	gtx.String(http.StatusOK, "User deleted successfully")
}
