package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type VehicleLogRepositoryInt interface {
	Create(*models.VehicleLog) (models.VehicleLog, error)
	Read(int) (models.VehicleLog, error)
	Update(*models.VehicleLog) (models.VehicleLog, error)
	Delete(int) error
	List() ([]models.VehicleLog, error)
}

type vehicleLogRepository struct {
	db *gorm.DB
}

func NewVehicleLogRepository(db *gorm.DB) VehicleLogRepositoryInt {
	return &vehicleLogRepository{db}
}

func (r *vehicleLogRepository) Create(item *models.VehicleLog) (models.VehicleLog, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return models.VehicleLog{}, err
	}
	return *item, nil
}

func (r *vehicleLogRepository) Delete(id int) error {
	err := r.db.Delete(&models.VehicleLog{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *vehicleLogRepository) List() ([]models.VehicleLog, error) {
	item := []models.VehicleLog{}
	var err error

	err = r.db.Find(&item).Error

	if err != nil {
		return []models.VehicleLog{}, err
	}
	return item, nil
}

func (r *vehicleLogRepository) Read(id int) (models.VehicleLog, error) {
	item := models.VehicleLog{}
	err := r.db.First(&item, id).Error
	if err != nil {
		return models.VehicleLog{}, err
	}
	return item, nil

}

func (r *vehicleLogRepository) Update(item *models.VehicleLog) (models.VehicleLog, error) {
	err := r.db.Save(&item).Error
	if err != nil {
		return models.VehicleLog{}, err
	}
	return *item, nil
}
