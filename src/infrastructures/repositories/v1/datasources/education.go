package datasources

import (
	"errors"
	"fmt"

	"github.com/prasetiyo28/go-solid-principle/src/domains"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EducationRepos struct {
	db *gorm.DB
}

func NewEducationRepos(db *gorm.DB) domains.EducationRepo {
	return &EducationRepos{
		db: db,
	}
}

func (e *EducationRepos) GetEducation(id string) (*domains.Education, error) {
	var education domains.Education
	var err error
	if err = e.db.Preload("Type").Where("id = ?", id).First(&education).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New(fmt.Sprintf("education with id %s is not found", id))
		}
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"id": err,
	}).Debug("id------")
	return &education, nil
}

func (e *EducationRepos) GetAllEducation() ([]*domains.Education, error) {
	var educations []*domains.Education
	if err := e.db.Preload("Type").Find(&educations).Error; err != nil {
		return nil, err
	}
	return educations, nil
}
