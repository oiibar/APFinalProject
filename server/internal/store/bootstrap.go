package store

import (
	"log"
	"os"
	"time"

	"final/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// EnsureDefaultAdmin creates a default admin user if the store is empty.
// This is only for demo/dev so the project can be tested immediately.
//
// You can override credentials with:
//
//	DEFAULT_ADMIN_EMAIL, DEFAULT_ADMIN_PASSWORD
func EnsureDefaultAdmin(st Store) {
	// If there is at least one user, do nothing.
	if len(st.ListUsers()) > 0 {
		return
	}

	email := os.Getenv("DEFAULT_ADMIN_EMAIL")
	if email == "" {
		email = "admin@example.com"
	}
	password := os.Getenv("DEFAULT_ADMIN_PASSWORD")
	if password == "" {
		password = "admin123"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to generate default admin password hash: %v", err)
		return
	}

	_, _ = st.CreateUser(models.User{
		Name:         "Admin",
		Email:        email,
		PasswordHash: string(hash),
		Role:         "admin",
		CreatedAt:    time.Now(),
	})
}
