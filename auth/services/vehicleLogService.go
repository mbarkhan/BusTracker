package services

import (
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type VehicleLogServiceInt interface {
	Create(*models.VehicleLog) (models.VehicleLog, error)
	Read(int) (models.VehicleLog, error)
	Update(*models.VehicleLog) (models.VehicleLog, error)
	Delete(int) error
	List() ([]models.VehicleLog, error)
}

type vehicleLogService struct {
	thisSrvc repositories.VehicleLogRepositoryInt
}

func (r *vehicleLogService) Create(item *models.VehicleLog) (models.VehicleLog, error) {
	return r.thisSrvc.Create(item)
}

func (r *vehicleLogService) Delete(itemId int) error {
	return r.thisSrvc.Delete(itemId)
}

func (r *vehicleLogService) List() ([]models.VehicleLog, error) {
	return r.thisSrvc.List()
}

func (r *vehicleLogService) Read(itemId int) (models.VehicleLog, error) {
	return r.thisSrvc.Read(itemId)
}

func (r *vehicleLogService) Update(item *models.VehicleLog) (models.VehicleLog, error) {
	return r.thisSrvc.Update(item)
}

func NewVehicleLogService(item repositories.VehicleLogRepositoryInt) VehicleLogServiceInt {
	return &vehicleLogService{thisSrvc: item}
}
