package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/1995parham-teaching/students/internal/request"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var ErrInvalidSigningMethod = errors.New("unexpected signing method")

const TokenHeaderLen = 2

type Auth struct {
	Key      []byte
	Username string
	Name     string
	Password string
	Logger   *zap.Logger
}

func (a Auth) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenHeader := strings.Fields(c.Request().Header.Get("Authorization"))

		if len(tokenHeader) != TokenHeaderLen {
			a.Logger.Error("invalid authorization header structure", zap.Strings("authorization", tokenHeader))

			return echo.ErrUnauthorized
		}

		tokenString := tokenHeader[1]

		token, err := jwt.ParseWithClaims(tokenString, new(jwt.RegisteredClaims),
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, ErrInvalidSigningMethod
				}

				return a.Key, nil
			})
		if err != nil {
			a.Logger.Error("token parsing failed", zap.Error(err))

			return echo.ErrUnauthorized
		}

		if !token.Valid {
			a.Logger.Error("token validation failed")

			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			a.Logger.Error("invalid token claims")

			return echo.ErrUnauthorized
		}

		if err := claims.Valid(); err != nil {
			a.Logger.Error("invalid claims", zap.Error(err))

			return echo.ErrUnauthorized
		}

		c.Set("name", claims.Subject)

		return next(c)
	}
}

func (a Auth) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		a.Logger.Error("invalid login request", zap.Error(err))

		return echo.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		a.Logger.Error("invalid login request", zap.Error(err))

		return echo.ErrBadRequest
	}

	if req.Username != a.Username || req.Password != a.Password {
		a.Logger.Error("unauthorized access",
			zap.String("username", req.Username),
			zap.String("password", req.Password),
		)

		return echo.ErrUnauthorized
	}

	claims := &jwt.RegisteredClaims{
		Issuer:    "students-summer-2022",
		Subject:   a.Name,
		Audience:  []string{"admin"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        a.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(a.Key)
	if err != nil {
		a.Logger.Error("signing token failed", zap.Error(err))

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, tokenString)
}

func (a Auth) Register(g *echo.Group) {
	g.POST("/login", a.Login)
}
