package usecase

import (
	"sort"

	"github.com/prasetiyo28/go-solid-principle/src/domains"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
)

type EducationUsecase struct {
	EduRepo domains.EducationRepo
}

func NewEducationUseCase(er domains.EducationRepo) domains.EducationUsecase {
	return &EducationUsecase{
		EduRepo: er,
	}
}

func (eu *EducationUsecase) GetEducation(id string) (*configs.ResponseSuccess, *configs.ResponseError) {
	education, err := eu.EduRepo.GetEducation(id)
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}

	return configs.Success(200, "OK", education), nil
}

func (eu *EducationUsecase) GetAllEducation() (*configs.ResponseSuccess, *configs.ResponseError) {
	educations, err := eu.EduRepo.GetAllEducation()
	if err != nil {
		return nil, configs.Failed(400, "FAILED", err.Error())
	}

	sort.SliceStable(educations, func(i, j int) bool {
		return educations[i].EducationName < educations[j].EducationName
	})

	return configs.Success(200, "OK", educations), nil
}
