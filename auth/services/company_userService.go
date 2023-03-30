package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
	"time"
)

type CompanyUserServiceInt interface {
	Create(*models.Company_Users) (models.Company_Users, error)
	Read(int) (models.Company_Users, error)
	Update(*models.Company_Users) (models.Company_Users, error)
	Delete(int) error
	List() ([]models.Company_Users, error)
	UpdateExpireDate(int, int, int) error
}

type companyUserService struct {
	cmpuser repositories.CompanyUserRepositoryInt
}

func (r *companyUserService) UpdateExpireDate(companyID int, UserID int, creditDays int) error {
	t := time.Now()
	t = t.AddDate(0, 0, creditDays)
	return r.cmpuser.UpdateExpireDate(companyID, UserID, t)
}

func (r *companyUserService) Create(cmpUser *models.Company_Users) (models.Company_Users, error) {
	return r.cmpuser.Create(cmpUser)
}

func (r *companyUserService) Delete(cmpUser int) error {
	return r.cmpuser.Delete(cmpUser)
}

func (r *companyUserService) List() ([]models.Company_Users, error) {
	return r.cmpuser.List()
}

func (r *companyUserService) Read(cmpUserId int) (models.Company_Users, error) {
	return r.cmpuser.Read(cmpUserId)
}

func (r *companyUserService) Update(cmpUser *models.Company_Users) (models.Company_Users, error) {
	return r.cmpuser.Update(cmpUser)
}

func NewCompanyUserService(icmpuser repositories.CompanyUserRepositoryInt) CompanyUserServiceInt {
	return &companyUserService{cmpuser: icmpuser}
}
