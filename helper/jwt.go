package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_SECRET"))

type jwtClaims struct {
	UserId   int
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(userId int, username string) (string, error) {
	claims := jwtClaims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_KEY)
}

func ValidateJWT(token string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWT_KEY, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*jwtClaims)
	if !ok || !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserId, nil
}
