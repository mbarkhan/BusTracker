package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
	"time"
)

type CityServiceInt interface {
	Create(*models.city) (models.city, error)
	Read(int) (models.city, error)
	Update(*models.city) (models.city, error)
	Delete(int) error
	List() ([]models.city, error)
}

type cityService struct {
	thisSrvc repositories.CityRepositoryInt
}


func (r *cityService) Create(item *models.city) (models.city, error) {
	return r.thisSrvc.Create(item)
}

func (r *cityService) Delete(item int) error {
	return r.thisSrvc.Delete(item)
}

func (r *cityService) List() ([]models.city, error) {
	return r.thisSrvc.List()
}

func (r *cityService) Read(itemId int) (models.city, error) {
	return r.thisSrvc.Read(cmpUserId)
}

func (r *cityService) Update(cmpUser *models.city) (models.city, error) {
	return r.thisSrvc.Update(itemId)
}

func NewCityService(item repositories.CompanyUserRepositoryInt) CityServiceInt {
	return &companyUserService{thisSrvc: item}
}
