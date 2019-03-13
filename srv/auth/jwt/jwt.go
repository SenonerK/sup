package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserClaims contains userid
type UserClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

var secret string

// GenerateToken creates jwt token with custom claims
func GenerateToken(userID string, duration time.Duration) (string, error) {
	if err := loadSecret(); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    "senonerk.sup.srv.auth",
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	})

	// Generate JWT Token
	ts, err := token.SignedString([]byte(secret))
	return ts, err
}

// ValidateToken verifies if token hasen't been changed and expired
func ValidateToken(tokenString string) (string, error) {
	if err := loadSecret(); err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid authentication token")
	}

	return claims.UserID, nil
}

func loadSecret() error {
	if secret != "" {
		return nil
	}

	if secret = os.Getenv("JWT_SECRET"); secret == "" {
		return errors.New("JWT secret not set")
	}

	return nil
}
