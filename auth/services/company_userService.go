package services

import "githubmbarkhanBusTracker/auth/models"

type CompanyUserServiceInt interface {
	Create(*models.Company_Users) (models.Company_Users, error)
	Read(int) (models.Company_Users, error)
	Update(*models.Company_Users) (models.Company_Users, error)
	Delete(int) error
	List() ([]models.Company_Users, error)
}

type companyUserService struct {
	cmpuser CompanyUserServiceInt
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

func NewCompanyUserService(icmpuser CompanyUserServiceInt) CompanyUserServiceInt {
	return &companyUserService{cmpuser: icmpuser}
}
