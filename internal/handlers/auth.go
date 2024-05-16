package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nervxz/msg-board/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

/*
Auth Sources https://www.youtube.com/watch?v=d4Y2DkKbxM0 9:48 - 44:30

	https://www.youtube.com/watch?v=97pIa_kYTqc
*/
func setupAuth(g *gin.RouterGroup, deps *dependencies) {
	h := &AuthHandler{deps: deps}
	h.bind(g)
}

type AuthHandler struct {
	deps *dependencies
}

func (h *AuthHandler) bind(g *gin.RouterGroup) {
	g.POST("/signup", h.signup)
	g.POST("/signin", h.signin)
	g.POST("/signout", h.signout)
}

func (h *AuthHandler) signup(gtx *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := gtx.BindJSON(&user); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Check if username already exists
	var exists bool
	err := h.deps.db.QueryRow("SELECT EXISTS(SELECT 1 FROM Users WHERE Username = $1)", user.Username).Scan(&exists)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to check username: %v", err)
		return
	}
	if exists {
		gtx.String(http.StatusConflict, "Username already taken")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to hash password: %v", err)
		return
	}

	// Insert the user into the database
	_, err = h.deps.db.Exec("INSERT INTO Users (Username, Password) VALUES ($1, $2)", user.Username, hashedPassword)
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to create user: %v", err)
		return
	}

	gtx.String(http.StatusOK, "User created successfully")
}

func (h *AuthHandler) signin(gtx *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := gtx.BindJSON(&user); err != nil {
		gtx.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Retrieve the user from the database
	var dbUser struct {
		UserID   int
		Username string
		Password string
	}
	err := h.deps.db.QueryRow("SELECT UserID, Username, Password FROM Users WHERE Username = $1", user.Username).Scan(&dbUser.UserID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			gtx.String(http.StatusUnauthorized, "InCorrect username or password")
			return
		}
		gtx.String(http.StatusInternalServerError, "Failed to query user: %v", err)
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		gtx.String(http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Create a session
	sessionToken := utils.GenerateToken()
	err = h.deps.redis.Set(context.Background(), sessionToken, dbUser.UserID, 24*time.Hour).Err()
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to create session: %v", err)
		return
	}
	// Set the session token as a cookie
	gtx.SetCookie("session_token", sessionToken, 3600*24, "/", "", false, true)

	// Return the session token
	gtx.JSON(http.StatusOK, gin.H{"token": sessionToken})
}

func (h *AuthHandler) signout(gtx *gin.Context) {
	token, err := gtx.Cookie("session_token")
	if err != nil {
		gtx.String(http.StatusBadRequest, "Missing authorization token")
		return
	}

	// Delete the session
	err = h.deps.redis.Del(context.Background(), token).Err()
	if err != nil {
		gtx.String(http.StatusInternalServerError, "Failed to delete session: %v", err)
		return
	}

	// Clear the session token cookie
	gtx.SetCookie("session_token", "", -1, "/", "", false, true)

	gtx.String(http.StatusOK, "Signed out successfully")
}
