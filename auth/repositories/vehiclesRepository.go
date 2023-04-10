package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type VehicleRepositoryInt interface {
	Create(*models.Vehicles) (models.Vehicles, error)
	Read(int) (models.Vehicles, error)
	Update(*models.Vehicles) (models.Vehicles, error)
	Delete(int) error
	List() ([]models.Vehicles, error)
	GetVehicleByMobile(string) (*models.Vehicles, error)
}

type vehicleRepository struct {
	db *gorm.DB
}

// Create implements VehicleRepositoryInt
func (r *vehicleRepository) Create(vehicle *models.Vehicles) (models.Vehicles, error) {
	err := r.db.Create(&vehicle).Error
	if err != nil {
		return models.Vehicles{}, err
	}
	return *vehicle, nil
}

// Delete implements VehicleRepositoryInt
func (r *vehicleRepository) Delete(id int) error {
	err := r.db.Delete(&models.Vehicles{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// List implements VehicleRepositoryInt
func (r *vehicleRepository) List() ([]models.Vehicles, error) {
	vehicles := []models.Vehicles{}
	var err error

	err = r.db.Find(&vehicles).Error

	if err != nil {
		return []models.Vehicles{}, err
	}
	return vehicles, nil
}

// Read implements VehicleRepositoryInt
func (r *vehicleRepository) Read(id int) (models.Vehicles, error) {
	vehicle := models.Vehicles{}
	err := r.db.First(&vehicle, id).Error
	if err != nil {
		return models.Vehicles{}, err
	}
	return vehicle, nil
}

// Update implements VehicleRepositoryInt
func (r *vehicleRepository) Update(vehicle *models.Vehicles) (models.Vehicles, error) {
	err := r.db.Save(&vehicle).Error
	if err != nil {
		return models.Vehicles{}, err
	}
	return *vehicle, nil
}

func NewVehicleRepository(db *gorm.DB) VehicleRepositoryInt {
	return &vehicleRepository{db}
}
func (r *vehicleRepository) GetVehicleByMobile(mobNum string) (*models.Vehicles, error) {

	vehicle := models.Vehicles{}
	error := r.db.Model(models.Vehicles{}).Where("mobile=?", mobNum).First(&vehicle).Error

	if error != nil {
		return &models.Vehicles{}, error
	}

	return &vehicle, nil
}
