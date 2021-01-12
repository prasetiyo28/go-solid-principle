package routers

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/applications/usecase/v1"
	"github.com/prasetiyo28/go-solid-principle/src/domains"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/middlewares"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/repositories/v1/datasources"
	controlers "github.com/prasetiyo28/go-solid-principle/src/interfaces/controllers/v1"
	"gorm.io/gorm"
)

type Handler struct {
	User       domains.UserController
	Middleware middlewares.IAuthMiddle
	// Location domains.LocController
}

func NewHandler(db *gorm.DB, rc *redis.Client) *Handler {

	return &Handler{
		User:       controlers.NewUserController(usecase.NewUserUseCase(datasources.NewUserRepo(db, rc))),
		Middleware: middlewares.NewAuthMiddleware(datasources.NewUserRepo(db, rc)),
	}
}

func (h *Handler) Register(v1 *echo.Group) {
	ev := v1.Group("/user")
	ev.GET("/profile/:id", h.User.GetUser, h.Middleware.MiddleGate())
	ev.POST("/profile", h.User.CreateUser)
	ev.PUT("/profile", h.User.UpdateUser)
	ev.DELETE("/profile/:id", h.User.DeleteUser)
	ev.POST("/login", h.User.Login)
}
