package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Auth Sources https://www.youtube.com/watch?v=d4Y2DkKbxM0 9:48 - 44:30
func setupAuth(g *gin.RouterGroup, deps *dependencies) {
    h := &AuthHandler{deps: deps}
    h.bind(g)
}

type AuthHandler struct {
    deps *dependencies
}

func (h *AuthHandler) bind(g *gin.RouterGroup) {
    g.POST("/signup", h.signup)

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

