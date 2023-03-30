package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type PassengerPathServiceInt interface {
	Create(*models.Passenger_Path) (models.Passenger_Path, error)
	Read(int) (models.Passenger_Path, error)
	Update(*models.Passenger_Path) (models.Passenger_Path, error)
	Delete(int) error
	List() ([]models.Passenger_Path, error)
}

type psnPathService struct {
	psn_path repositories.PsnPathRepositoryInt
}

func (r *psnPathService) Create(psnPath *models.Passenger_Path) (models.Passenger_Path, error) {
	return r.psn_path.Create(psnPath)
}

func (r *psnPathService) Delete(psnPath int) error {
	return r.psn_path.Delete(psnPath)
}

func (r *psnPathService) List() ([]models.Passenger_Path, error) {
	return r.psn_path.List()
}

func (r *psnPathService) Read(cmpId int) (models.Passenger_Path, error) {
	return r.psn_path.Read(cmpId)
}

func (r *psnPathService) Update(psnPath *models.Passenger_Path) (models.Passenger_Path, error) {
	return r.psn_path.Update(psnPath)
}

func NewPsnPathService(ipsn repositories.PsnPathRepositoryInt) PassengerPathServiceInt {
	return &psnPathService{psn_path: ipsn}
}
