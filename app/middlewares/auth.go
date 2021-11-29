package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"presence-app-backend/controllers"
	"time"
)

type JwtClaims struct {
	UserId int
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (config *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: []byte(config.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(
			func(err error, c echo.Context) error {
				return controllers.NewErrorResponse(c, http.StatusForbidden, err)
			},
		),
	}
}

func (config *ConfigJWT) GenerateToken(userId int) (string, error) {
	claims := JwtClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(config.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(config.SecretJWT))
	return token, err
}
