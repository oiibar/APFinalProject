package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secretus")

type contextKey string

const ContextUserID = contextKey("user_id")

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

		var userID int
		switch v := userIDf.(type) {
		case float64:
			userID = int(v)
		case string:
			userID, _ = strconv.Atoi(v)
		}
		ctx := context.WithValue(r.Context(), ContextUserID, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
