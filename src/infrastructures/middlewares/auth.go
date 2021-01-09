package middlewares

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddle struct {
	DB *gorm.DB
}

type IAuthMiddle interface {
	MiddleGate() echo.MiddlewareFunc
}

func NewAuthMiddleware(db *gorm.DB) IAuthMiddle {
	return &AuthMiddle{
		DB: db,
	}
}

func (am *AuthMiddle) MiddleGate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()
			res.Header().Set("X-ORIGIN", "up-to-the-moon")
			return next(c)
		}
	}
}