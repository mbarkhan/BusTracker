package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type ProvinceServiceInt interface {
	Create(*models.Province) (models.Province, error)
	Read(int) (models.Province, error)
	Update(*models.Province) (models.Province, error)
	Delete(int) error
	List() ([]models.Province, error)
}

type provinceService struct {
	thisSrvc repositories.ProvinceRepositoryInt
}

func (r *provinceService) Create(item *models.Province) (models.Province, error) {
	return r.thisSrvc.Create(item)
}

func (r *provinceService) Delete(itemId int) error {
	return r.thisSrvc.Delete(itemId)
}

func (r *provinceService) List() ([]models.Province, error) {
	return r.thisSrvc.List()
}

func (r *provinceService) Read(itemId int) (models.Province, error) {
	return r.thisSrvc.Read(itemId)
}

func (r *rovinceService) Update(item *models.Province) (models.Province, error) {
	return r.thisSrvc.Update(item)
}

func NewrovinceService(item repositories.ProvinceRepositoryInt) ProvinceServiceInt {
	return &ProvinceService{thisSrvc: item}
}
