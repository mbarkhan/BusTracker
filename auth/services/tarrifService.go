package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type TarrifServiceInt interface {
	Create(*models.Tariff) (models.Tariff, error)
	Read(int) (models.Tariff, error)
	Update(*models.Tariff) (models.Tariff, error)
	Delete(int) error
	List() ([]models.Tariff, error)
}

type tarrifService struct {
	thisSrvc repositories.TarrifsRepositoryInt
}

func (r *tarrifService) Create(item *models.Tariff) (models.Tariff, error) {
	return r.thisSrvc.Create(item)
}

func (r *tarrifService) Delete(itemId int) error {
	return r.thisSrvc.Delete(itemId)
}

func (r *tarrifService) List() ([]models.Tariff, error) {
	return r.thisSrvc.List()
}

func (r *tarrifService) Read(itemId int) (models.Tariff, error) {
	return r.thisSrvc.Read(itemId)
}

func (r *tarrifService) Update(item *models.Tariff) (models.Tariff, error) {
	return r.thisSrvc.Update(item)
}

func NewTarrifService(item repositories.TarrifsRepositoryInt) TarrifServiceInt {
	return &tarrifService{thisSrvc: item}
}
