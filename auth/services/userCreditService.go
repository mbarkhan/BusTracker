package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type UserCreditServiceInt interface {
	Create(*models.UserCredit) (models.UserCredit, error)
	Read(uint) ([]models.UserCredit, error)
}

type userCreditService struct {
	userCredit repositories.UserCreditRepositoryInt
}

func NewUserCreditService(iuser repositories.UserCreditRepositoryInt) UserCreditServiceInt {
	return &userCreditService{userCredit: iuser}
}

func (r *userCreditService) Create(userCredit *models.UserCredit) (models.UserCredit, error) {
	var err error
	if err != nil {
		return models.UserCredit{}, err
	}
	return r.userCredit.Create(userCredit)
}

func (r *userCreditService) Read(userId uint) ([]models.UserCredit, error) {
	return r.userCredit.Read(userId)
}
