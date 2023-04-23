package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type VehicleServiceInt interface {
	Create(*models.Vehicle) (models.Vehicle, error)
	Read(int) (models.Vehicle, error)
	Update(*models.Vehicle) (models.Vehicle, error)
	Delete(int) error
	List() ([]models.Vehicle, error)
}

type vehicleService struct {
	vehicle repositories.VehicleRepositoryInt
}

func (r *vehicleService) Create(vhcl *models.Vehicle) (models.Vehicle, error) {
	return r.vehicle.Create(vhcl)
}

func (r *vehicleService) Delete(vhcl int) error {
	return r.vehicle.Delete(vhcl)
}

func (r *vehicleService) List() ([]models.Vehicle, error) {
	return r.vehicle.List()
}

func (r *vehicleService) Read(cmpId int) (models.Vehicle, error) {
	return r.vehicle.Read(cmpId)
}

func (r *vehicleService) Update(vhcl *models.Vehicle) (models.Vehicle, error) {
	return r.vehicle.Update(vhcl)
}

func NewvehicleService(ivcl repositories.VehicleRepositoryInt) VehicleServiceInt {
	return &vehicleService{vehicle: ivcl}
}
