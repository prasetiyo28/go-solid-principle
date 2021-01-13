package domains

import (
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
)

func (Education) TableName() string {
	return "educations"
}

type Education struct {
	ID            int32  `gorm:"column:id;PRIMARY_KEY" json:"id"`
	EducationName string `gorm:"column:education_name;" json:"education_name"`
	TypeID        int32  `gorm:"column:type_id;" json:"type_id"`
	Type          *Type  `gorm:"foreignkey:TypeID" json:"type"`
}

type EducationRepo interface {
	GetEducation(string) (*Education, error)
	GetAllEducation() ([]*Education, error)
}

type EducationUsecase interface {
	GetEducation(string) (*configs.ResponseSuccess, *configs.ResponseError)
	GetAllEducation() (*configs.ResponseSuccess, *configs.ResponseError)
}

type EducationController interface {
	GetEducation(echo.Context) error
	GetAllEducation(echo.Context) error
}

type Educations struct {
	Educations []Education `json:"educations"`
}

func (e *Education) Bind(c echo.Context) error {
	if err := c.Bind(e); err != nil {
		return err
	}
	return nil
}
