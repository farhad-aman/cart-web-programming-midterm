package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtSecretKey = []byte("secret")

type JwtClaims struct {
	jwt.StandardClaims
	UserID uint
}

func GenerateToken(UserID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JwtClaims{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecretKey)

	return tokenString, err
}

func ParseToken(tokenStr string) (*JwtClaims, error) {
	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
