package services

import "githubmbarkhanBusTracker/auth/models"

type VehicleServiceInt interface {
	Create(*models.Vehicles) (models.Vehicles, error)
	Read(int) (models.Vehicles, error)
	Update(*models.Vehicles) (models.Vehicles, error)
	Delete(int) error
	List() ([]models.Vehicles, error)
}

type vehicleService struct {
	vehicle VehicleServiceInt
}

func (r *vehicleService) Create(vhcl *models.Vehicles) (models.Vehicles, error) {
	return r.vehicle.Create(vhcl)
}

func (r *vehicleService) Delete(vhcl int) error {
	return r.vehicle.Delete(vhcl)
}

func (r *vehicleService) List() ([]models.Vehicles, error) {
	return r.vehicle.List()
}

func (r *vehicleService) Read(cmpId int) (models.Vehicles, error) {
	return r.vehicle.Read(cmpId)
}

func (r *vehicleService) Update(vhcl *models.Vehicles) (models.Vehicles, error) {
	return r.vehicle.Update(vhcl)
}

func NewvehicleService(ivcl VehicleServiceInt) VehicleServiceInt {
	return &vehicleService{vehicle: ivcl}
}
