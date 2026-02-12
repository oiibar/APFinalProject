package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secretus")

type contextKey string

const ContextUserID = contextKey("user_id")
const ContextUserRole = contextKey("user_role")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract Bearer Token
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			next.ServeHTTP(w, r)
			return
		}

		// Parse & Validate JWT
		tokStr := parts[1]
		token, err := jwt.Parse(tokStr, func(token *jwt.Token) (interface{}, error) {
			// Enforce expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return JwtKey, nil
		})
		if err != nil || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}
		userIDf, ok := claims["user_id"]
		if !ok {
			next.ServeHTTP(w, r)
			return
		}
		role, _ := claims["role"].(string)

		var userID int
		switch v := userIDf.(type) {
		case float64:
			userID = int(v)
		case string:
			userID, _ = strconv.Atoi(v)
		}
		ctx := context.WithValue(r.Context(), ContextUserID, userID)
		ctx = context.WithValue(ctx, ContextUserRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) (int, bool) {
	uid := r.Context().Value(ContextUserID)
	if uid == nil {
		return 0, false
	}
	id, ok := uid.(int)
	if !ok || id <= 0 {
		return 0, false
	}
	return id, true
}

func GetUserRole(r *http.Request) string {
	role, _ := r.Context().Value(ContextUserRole).(string)
	return role
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := GetUserID(r); !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := GetUserID(r); !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}
		if strings.ToLower(GetUserRole(r)) != "admin" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "forbidden"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
