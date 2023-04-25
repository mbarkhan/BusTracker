package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type TariffServiceInt interface {
	Create(*models.Tariff) (models.Tariff, error)
	Read(int) (models.Tariff, error)
	Update(*models.Tariff) (models.Tariff, error)
	Delete(int) error
	List() ([]models.Tariff, error)
}

type tariffService struct {
	thisSrvc repositories.TariffRepositoryInt
}

func (r *tariffService) Create(item *models.Tariff) (models.Tariff, error) {
	return r.thisSrvc.Create(item)
}

func (r *tariffService) Delete(itemId int) error {
	return r.thisSrvc.Delete(itemId)
}

func (r *tariffService) List() ([]models.Tariff, error) {
	return r.thisSrvc.List()
}

func (r *tariffService) Read(itemId int) (models.Tariff, error) {
	return r.thisSrvc.Read(itemId)
}

func (r *tariffService) Update(item *models.Tariff) (models.Tariff, error) {
	return r.thisSrvc.Update(item)
}

func NewTarrifService(item repositories.TariffRepositoryInt) TariffServiceInt {
	return &tariffService{thisSrvc: item}
}
