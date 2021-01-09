package domains

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
	"gorm.io/gorm"
)

func (User) TableName() string {
	return "user"
}

type User struct {
	UserID   string `gorm:"column:user_id;" json:"user_id"`
	Email    string `gorm:"column:email" json:"email"`
	Address  string `gorm:"column:address" json:"address"`
	Password string `gorm:"column:password" json:"password"`
}

type UserRepo interface {
	GetUser(string) (*User, error)
	CreateUser(User) (*User, error)
	UpdateUser(User) (*User, error)
	DeleteUser(string) (*string, error)
	GetUserByEmail(User) (*User, error)
}


 
type UserUseCase interface {
	GetUser(interface{}) (*configs.ResponseSuccess, *configs.ResponseError)
	CreateUser(User) (*configs.ResponseSuccess, *configs.ResponseError)
	UpdateUser(User) (*configs.ResponseSuccess, *configs.ResponseError)
	DeleteUser(interface{}) (*configs.ResponseSuccess, *configs.ResponseError)
	Login(User) (*configs.ResponseSuccess, *configs.ResponseError)
}

type UserController interface {
	GetUser(echo.Context) error
	CreateUser(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
	Login(echo.Context) error
}

func (e *User) Bind(c echo.Context) error {
	if err := c.Bind(e); err != nil {
		return err
	}
	return nil
}

func (e *User) BeforeCreate(tx *gorm.DB) (err error) {
	var user User
	var count int64
	if e.UserID == "" {
		e.UserID = uuid.New().String()
	}

	for {
		if err := tx.Model(&user).Where("user_id = ?", e.UserID).Count(&count).Error; err != nil {
			return err
		}
		if count < 1 {
			break
		}
		e.UserID = uuid.New().String()
	}

	return
}
