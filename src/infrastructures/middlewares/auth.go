package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/domains"
)

type AuthMiddle struct {
	// DB *gorm.DB  // repo user
	UsRep domains.UserRepo
}

type IAuthMiddle interface {
	MiddleGate() echo.MiddlewareFunc
}

func NewAuthMiddleware(UsRep domains.UserRepo) IAuthMiddle { //repo user
	return &AuthMiddle{
		UsRep: UsRep, //repo user
	}
}

func (UsRep *AuthMiddle) MiddleGate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()

			res.Header().Set("X-ORIGIN", "up-to-the-moon")
			req := c.Request()
			if req == nil {
				panic("token needed")
			}
			token := req.Header.Get("auth")
			UsRep.UsRep.GetToken(token)
			return next(c)
		}
	}
}
