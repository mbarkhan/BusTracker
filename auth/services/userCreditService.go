package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type UserCreditServiceInt interface {
	Create(*models.User_Credit) (models.User_Credit, error)
	Read(int) ([]models.User_Credit, error)
}

type userCreditService struct {
	userCredit repositories.UserCreditRepositoryInt
}

func NewUserCreditService(iuser repositories.UserCreditRepositoryInt) UserCreditServiceInt {
	return &userCreditService{userCredit: iuser}
}

func (r *userCreditService) Create(userCredit *models.User_Credit) (models.User_Credit, error) {
	var err error
	if err != nil {
		return models.User_Credit{}, err
	}
	return r.userCredit.Create(userCredit)
}

func (r *userCreditService) Read(userId int) ([]models.User_Credit, error) {
	return r.userCredit.Read(userId)
}
