package usecase

import (
	"fmt"
	"time"

	"github.com/prasetiyo28/go-solid-principle/src/domains"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UsRepo domains.UserRepo
}

func NewUserUseCase(er domains.UserRepo) domains.UserUseCase {
	return &UserUseCase{
		UsRepo: er,
	}
}

func (uu *UserUseCase) GetUser(id interface{}) (*configs.ResponseSuccess, *configs.ResponseError) {
	s, ok := id.(string)
	if !ok {
		return nil, configs.Failed(400, "FAILED", "id must be a string")
	}

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	// user, err := uu.UsRepo.GetUser(s)
	// if err != nil {
	// 	return nil, configs.Failed(400, "FAILED", err.Error())
	// }

	user, err := uu.UsRepo.GetToken(s)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}

	return configs.Success(200, "OK", user), nil
}

func (uu *UserUseCase) CreateUser(us domains.User) (*configs.ResponseSuccess, *configs.ResponseError) {

	email, err := uu.UsRepo.GetUserByEmail(us)
	if email != nil {
		return nil, configs.Failed(400, "FAILED REGISTER", "Email Already Exist")
	}

	const DefaultCost int = 10
	beforHash := []byte(us.Password)
	hash, err := bcrypt.GenerateFromPassword(beforHash, DefaultCost)
	if err != nil {
		return nil, configs.Failed(400, "FAILED HASH", err.Error())
	}

	us.Password = string(hash)
	user, err := uu.UsRepo.CreateUser(us)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", user), nil
}

func (uu *UserUseCase) UpdateUser(us domains.User) (*configs.ResponseSuccess, *configs.ResponseError) {
	if us.Password != "" {
		const DefaultCost int = 10
		beforHash := []byte(us.Password)
		hash, err := bcrypt.GenerateFromPassword(beforHash, DefaultCost)
		if err != nil {
			return nil, configs.Failed(400, "FAILED HASH", err.Error())
		}
		us.Password = string(hash)
	}

	user, err := uu.UsRepo.UpdateUser(us)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", user), nil
}

func (uu *UserUseCase) DeleteUser(id interface{}) (*configs.ResponseSuccess, *configs.ResponseError) { //string
	s, ok := id.(string)
	if !ok {
		return nil, configs.Failed(400, "FAILED", "id must be a string")
	}
	user, errCheck := uu.UsRepo.GetUser(s)
	if errCheck != nil {
		return nil, configs.Failed(400, "FAILED", errCheck.Error())
	}

	id, errDel := uu.UsRepo.DeleteUser(user.UserID)
	if errDel != nil {
		return nil, configs.Failed(400, "FAILED", errDel.Error())
	}
	return configs.Success(200, "OK", user), nil
}

func (uu *UserUseCase) Login(us domains.User) (*configs.ResponseSuccess, *configs.ResponseError) {
	user, err := uu.UsRepo.GetUserByEmail(us)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	comparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(us.Password))
	if comparePass != nil {
		return nil, configs.Failed(400, "FAILED", "Password Didn't match")
	}

	currentTime := time.Now().String()
	beforHash := []byte(currentTime)
	hash, err := bcrypt.GenerateFromPassword(beforHash, 10)
	if err != nil {
		return nil, configs.Failed(400, "FAILED HASH", err.Error())
	}
	token := string(hash)

	errToken := uu.UsRepo.SetToken(token, user)
	if errToken != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}
	return configs.Success(200, "OK", token), nil
}
