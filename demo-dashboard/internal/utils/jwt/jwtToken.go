package jwtToken

import (
	"demo-dashboard/internal/conf"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenToken(subject string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    "eeo",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(conf.Jwt.Expired))),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signToken, err := token.SignedString([]byte(conf.Jwt.Secret))

	return signToken, err
}
