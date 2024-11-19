package jwt

import (
	"demo-base/internal/conf"
	"demo-base/internal/utils/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//生成JWT，存放在localStorage中，每次请求的时候，放在header Authorization中，payload中包含用户信息，过期时间，签发时间，签发人，校验JWT

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}

const (
	// TokenExpireDuration token 过期时间
	TokenExpireDuration = 24 * time.Hour
	keyword             = "zU9uB5fR7tX1tZ3bV4kG0xT2lY5eR2sP"
)

func NewJWT(username string) string {
	signature, err := newToken(username).SignedString([]byte(keyword))
	if err != nil {
		logger.Panic(err.Error())
	}

	return signature
}

func newToken(username string) *jwt.Token {
	jwtClaim := &JWTClaim{
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(conf.Jwt.Expired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    conf.Issuer,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
}

func ParseToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(keyword), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token invalid")
	}
	if claims, ok := token.Claims.(*JWTClaim); ok {
		return claims, nil
	}

	return nil, nil
}
