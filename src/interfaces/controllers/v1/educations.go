package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/domains"
)

type EducationController struct {
	UseCase domains.EducationUsecase
}

func NewEducationController(uc domains.EducationUsecase) domains.EducationController {
	return &EducationController{
		UseCase: uc,
	}
}

func (uc *EducationController) GetEducation(c echo.Context) error {
	result, err := uc.UseCase.GetEducation(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(result.Status, result)
}

func (uc *EducationController) GetAllEducation(c echo.Context) error {
	result, err := uc.UseCase.GetAllEducation()
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(result.Status, result)
}
