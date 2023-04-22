package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type TarrifServiceInt interface {
	Create(*models.Tarrif) (models.Tarrif, error)
	Read(int) (models.Tarrif, error)
	Update(*models.Tarrif) (models.Tarrif, error)
	Delete(int) error
	List() ([]models.Tarrif, error)
}

type tarrifService struct {
	thisSrvc repositories.RepositoryInt
}

func (r *tarrifService) Create(item *models.Tarrif) (models.Tarrif, error) {
	return r.thisSrvc.Create(item)
}

func (r *tarrifService) Delete(itemId int) error {
	return r.thisSrvc.Delete(itemId)
}

func (r *tarrifService) List() ([]models.Tarrif, error) {
	return r.thisSrvc.List()
}

func (r *tarrifService) Read(itemId int) (models.Tarrif, error) {
	return r.thisSrvc.Read(itemId)
}

func (r *tarrifService) Update(item *models.Tarrif) (models.Tarrif, error) {
	return r.thisSrvc.Update(item)
}

func NewTarrifService(item repositories.TarrifRepositoryInt) TarrifServiceInt {
	return &tarrifService{thisSrvc: item}
}
