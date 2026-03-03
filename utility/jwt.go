package utility

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPurpose string

const (
	JwtPurposeAccess  JwtPurpose = "Access"
	JwtPurposeRefresh JwtPurpose = "Refresh"
)

type JwtClaims struct {
	Purpose JwtPurpose `json:"purpose"`
	jwt.RegisteredClaims
}

func GenToken(userId int, purpose JwtPurpose, expireTime time.Duration) (string, error) {
	claims := JwtClaims{purpose, jwt.RegisteredClaims{
		Issuer:    "https://stavi.tw",
		Subject:   fmt.Sprintf("%d", userId),
		Audience:  jwt.ClaimStrings{"https://api.stavi.tw"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JwtSecret)
}

func ParseToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (any, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
