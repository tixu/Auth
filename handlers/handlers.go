package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaim struct {
	Email string `json:Email`
	Roles string `json:Roles`
	jwt.StandardClaims
}
type VersionResponse struct {
	Version string `json:"version"`
}

type versionHandler struct {
	version string
}

func (h *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version: h.version,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func VersionHandler(version string) http.Handler {
	return &versionHandler{
		version: version,
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type loginHandler struct {
	secret string
	users  Users
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	user, ok := h.users[username]
	if !ok {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
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
		Subject:   user.Username,
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

func LoginHandler(secret string, users Users) http.Handler {
	return &loginHandler{
		secret: secret,
		users:  users,
	}
}

type User struct {
	Username     string
	PasswordHash string
	Role         string
	Email        string
}

type Users map[string]User

var DB = Users{
	"user": User{
		Username: "user",
		// bcrypt has for "password"
		PasswordHash: "$2a$10$KgFhp4HAaBCRAYbFp5XYUOKrbO90yrpUQte4eyafk4Tu6mnZcNWiK",
		Email:        "user@example.com",
		Role:         "wtfd",
	},
}
