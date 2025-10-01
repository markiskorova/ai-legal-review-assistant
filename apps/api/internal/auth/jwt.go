package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Sign(userID int64) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(key)
}
