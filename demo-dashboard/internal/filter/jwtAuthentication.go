package filter

import (
	"demo-dashboard/internal/conf"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthentication() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    []byte(conf.Jwt.Secret),
		},
		ContextKey: "username",
		Claims:     jwt.RegisteredClaims{},
	})
}
