package datasources

import (
	"errors"
	"fmt"

	"github.com/prasetiyo28/go-solid-principle/src/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepos struct {
	db *gorm.DB
	// cache *redis.Client
}

func NewUserRepo(db *gorm.DB) domains.UserRepo {
	return &UserRepos{
		db: db,
		// cache: cache,
	}
}

func (r *UserRepos) GetUser(id string) (*domains.User, error) {
	var user domains.User
	var err error
	if err = r.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New(fmt.Sprintf("user with id %s is not found", id))
		}
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"id": err,
	}).Debug("id------")
	return &user, nil
}

func (r *UserRepos) CreateUser(us domains.User) (*domains.User, error) {
	var user domains.User
	result := r.db.Omit(clause.Associations).Create(&us)

	if result.Error != nil {
		return nil, result.Error
	}
	user = us
	logrus.WithFields(logrus.Fields{
		"us": us,
	}).Debug("ev-----")
	return &user, nil
}

func (r *UserRepos) UpdateUser(us domains.User) (*domains.User, error) {
	var user domains.User
	result := r.db.Model(&user).Where("user_id = ?", us.UserID).Updates(us)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New(fmt.Sprintf("there isn't row affected"))
	}
	user = us
	return &user, nil
}

func (r *UserRepos) DeleteUser(id string) (*string, error) {
	if err := r.db.Where("user_id = ?", id).Delete(&domains.User{}).Error; err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *UserRepos) GetUserByEmail(us domains.User) (*domains.User, error) {
	var user domains.User
	var err error
	if err = r.db.Where("email = ?", us.Email).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New(fmt.Sprintf("email %s is not found", us.Email))
		}
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"email": err,
	}).Debug("email------")
	return &user, nil
}
