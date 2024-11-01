package jwt

import (
	"crypto/rand"
	"demo-base/internal/conf"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//生成JWT，存放在localStorage中，每次请求的时候，放在header Authorization中，payload中包含用户信息，过期时间，签发时间，签发人，校验JWT

type JWT struct {
	Key       []byte //rand.Read() 生成的随机字符串
	Token     *jwt.Token
	Signature string
}

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}

const (
	// TokenExpireDuration token 过期时间
	TokenExpireDuration = 24 * time.Hour
)

func NewJWT(username string) *JWT {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	signature, err := NewToken(username).SignedString(key)
	if err != nil {
		panic(err)
	}

	return &JWT{
		Key:       key,
		Token:     NewToken(username),
		Signature: signature,
	}
}

func NewToken(username string) *jwt.Token {
	jwtClaim := &JWTClaim{
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(conf.Jwt.Expired * int(time.Second)))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    conf.Issuer,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
}

func (j *JWT) GenerateToken(username string) (string, error) {
	return j.Token.SignedString(j.Key)
}

func (j *JWT) ParseToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token invalid")
	}
	if claims, ok := token.Claims.(*JWTClaim); ok {
		return claims, nil
	}

	return nil, nil
}
