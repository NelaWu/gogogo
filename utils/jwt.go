package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("123456")

// GenerateToken 生成 JWT token
func GenerateToken(userID uint) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken 驗證 JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || token.Valid == false {
		return nil, err
	}
	return claims, nil
}
