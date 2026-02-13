package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"final/internal/models"
	"final/internal/store"
)

var JwtKey = []byte("secretus")

type AuthHandler struct {
	Store store.Store
}

type signupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	if req.Name == "" || req.Email == "" || len(req.Password) < 6 {
		writeError(w, http.StatusBadRequest, "name, email and password (min 6 chars) are required")
		return
	}

	if _, ok := h.Store.GetUserByEmail(req.Email); ok {
		writeError(w, http.StatusConflict, "email already registered")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	u := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "user",
		CreatedAt:    time.Now(),
	}

	created, _ := h.Store.CreateUser(u)
	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"id":    created.ID,
		"name":  created.Name,
		"email": created.Email,
		"role":  created.Role,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" || req.Password == "" {
		writeError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	u, ok := h.Store.GetUserByEmail(req.Email)
	if !ok {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	exp := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"role":    u.Role,
		"exp":     exp.Unix(),
	})
	ss, err := token.SignedString(JwtKey)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not sign token")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"access_token": ss,
		"token_type":   "bearer",
		"expires_in":   int(time.Until(exp).Seconds()),
		"user": map[string]interface{}{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
			"role":  u.Role,
		},
	})
}
