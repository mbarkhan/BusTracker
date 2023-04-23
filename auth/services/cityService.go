package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type CityServiceInt interface {
	Create(*models.City) (models.City, error)
	Read(int) (models.City, error)
	Update(*models.City) (models.City, error)
	Delete(int) error
	List() ([]models.City, error)
}

type cityService struct {
	thisSrvc repositories.CityRepositoryInt
}

func (r *cityService) Create(item *models.City) (models.City, error) {
	return r.thisSrvc.Create(item)
}

func (r *cityService) Delete(item int) error {
	return r.thisSrvc.Delete(item)
}

func (r *cityService) List() ([]models.City, error) {
	return r.thisSrvc.List()
}

func (r *cityService) Read(itemId int) (models.City, error) {
	return r.thisSrvc.Read(itemId)
}

func (r *cityService) Update(itemId *models.City) (models.City, error) {
	return r.thisSrvc.Update(itemId)
}

func NewCityService(item repositories.CityRepositoryInt) CityServiceInt {
	return &cityService{thisSrvc: item}
}
