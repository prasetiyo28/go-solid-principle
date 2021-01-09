package controlers

import (
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/domains"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
)

type UserController struct {
	UseCase domains.UserUseCase
}

func NewUserController(uc domains.UserUseCase) domains.UserController {
	return &UserController{
		UseCase: uc,
	}
}

func (ec *UserController) GetUser(c echo.Context) error {
	result, err := ec.UseCase.GetUser(c.Param("id"))
	if err != nil {
		// return errors.New("ERROR")
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var user domains.User
	if err := user.Bind(c); err != nil {
		failed := configs.Failed(400, "FAILED", err.Error())
		return c.JSON(failed.Status, failed)
	}
	result, err := uc.UseCase.CreateUser(user)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	var user domains.User
	if err := user.Bind(c); err != nil {
		failed := configs.Failed(400, "FAILED", err.Error())
		return c.JSON(failed.Status, failed)
	}
	result, err := uc.UseCase.UpdateUser(user)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(result.Status, result)
}
func (uc *UserController) DeleteUser(c echo.Context) error {
	result, err := uc.UseCase.DeleteUser(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(result.Status, result)
}

func (uc *UserController) Login(c echo.Context) error {
	var user domains.User
	if err := user.Bind(c); err != nil {
		failed := configs.Failed(400, "FAILED", err.Error())
		return c.JSON(failed.Status, failed)
	}
	result, err := uc.UseCase.Login(user)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(result.Status, result)
}
