package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
)

type ErrorMiddleware struct {
}

type IErrorMiddleware interface {
	HttpErrorHandler(error, echo.Context)
}

func NewErrorMiddleware() IErrorMiddleware {
	return &ErrorMiddleware{}
}

func (em *ErrorMiddleware) HttpErrorHandler(e error, c echo.Context) {
	err := configs.Failed(400, "ERROR", e.Error())
	c.JSON(err.Status, err)
}
