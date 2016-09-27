package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tixu/Auth/users"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaim struct {
	Email string `json:Email`
	Roles string `json:Roles`
	jwt.StandardClaims
}

type loginHandler struct {
	secret      string
	userservice users.UserService
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Client creates a connection to the services.
type Client interface {
	Connect() Session
}

// Session represents authenticable connection to the services.
type Session interface {
	UserService() users.User
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	user, err := h.userservice.GetUser(username)
	if err != nil {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	standartclaims := jwt.StandardClaims{
		Issuer:    "auth.service",
		Audience:  "tixu",
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   user.Name,
	}

	claims := CustomClaim{
		user.Email,
		user.Role,
		standartclaims,
	}

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(h.secret))
	if err != nil {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		Token: tokenString,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(secret string, userService users.UserService) http.Handler {
	return &loginHandler{
		secret:      secret,
		userservice: userService,
	}
}
