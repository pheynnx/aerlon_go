package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Password string `json:"password"`
	Pin      string `json:"pin"`
	jwt.RegisteredClaims
}

func CreateAdminJWT(password, pin string) (string, error) {
	claims := JWTClaims{
		password,
		pin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func VerifyAdminJWT(token string) bool {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false
	}

	if claims, ok := parsedToken.Claims.(*JWTClaims); ok && parsedToken.Valid {
		if claims.Password == os.Getenv("ADMIN_PASSWORD") && claims.Pin == os.Getenv("ADMIN_PIN") {
			return true
		}
	}

	return false
}
