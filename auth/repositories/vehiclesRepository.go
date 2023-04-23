package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type VehicleRepositoryInt interface {
	Create(*models.Vehicle) (models.Vehicle, error)
	Read(int) (models.Vehicle, error)
	Update(*models.Vehicle) (models.Vehicle, error)
	Delete(int) error
	List() ([]models.Vehicle, error)
}

type vehicleRepository struct {
	db *gorm.DB
}

// Create implements VehicleRepositoryInt
func (r *vehicleRepository) Create(vehicle *models.Vehicle) (models.Vehicle, error) {
	err := r.db.Create(&vehicle).Error
	if err != nil {
		return models.Vehicle{}, err
	}
	return *vehicle, nil
}

// Delete implements VehicleRepositoryInt
func (r *vehicleRepository) Delete(id int) error {
	err := r.db.Delete(&models.Vehicle{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// List implements VehicleRepositoryInt
func (r *vehicleRepository) List() ([]models.Vehicle, error) {
	vehicles := []models.Vehicle{}
	var err error

	err = r.db.Find(&vehicles).Error

	if err != nil {
		return []models.Vehicle{}, err
	}
	return vehicles, nil
}

// Read implements VehicleRepositoryInt
func (r *vehicleRepository) Read(id int) (models.Vehicle, error) {
	vehicle := models.Vehicle{}
	err := r.db.First(&vehicle, id).Error
	if err != nil {
		return models.Vehicle{}, err
	}
	return vehicle, nil
}

// Update implements VehicleRepositoryInt
func (r *vehicleRepository) Update(vehicle *models.Vehicle) (models.Vehicle, error) {
	err := r.db.Save(&vehicle).Error
	if err != nil {
		return models.Vehicle{}, err
	}
	return *vehicle, nil
}

func NewVehicleRepository(db *gorm.DB) VehicleRepositoryInt {
	return &vehicleRepository{db}
}
