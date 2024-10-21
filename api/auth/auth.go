package auth

import (
	"fmt"

	"github.com/realjv3/notifs/users"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB, r *gin.Engine) *Service {

	// Setup JWT middleware that parses the token
	// and injects the claims into gin's context
	r.Use(jwtMiddleware())

	return &Service{db}
}

// Authenticate authenticates a user by email
func (s Service) Authenticate(email string) (string, error) {
	var u users.User
	err := s.db.Get(&u, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		// TODO: log error and then obfuscate error in output - this is insecure
		return "", fmt.Errorf("failed to get user: %s", err)
	}

	tokenString, err := createJWT(email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %s", err)
	}

	return tokenString, nil
}
