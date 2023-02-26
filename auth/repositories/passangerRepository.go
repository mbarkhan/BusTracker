package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type PassengerRepositoryInt interface {
	Create(*models.Passenger_Path) (models.Passenger_Path, error)
	Read(int) (models.Passenger_Path, error)
	Update(*models.Passenger_Path) (models.Passenger_Path, error)
	Delete(int) error
	List() ([]models.Passenger_Path, error)
}

type passengerRepository struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) PassengerRepositoryInt {
	return &passengerRepository{db}
}

// Create implements PassengerRepositoryInt
func (r *passengerRepository) Create(passenger *models.Passenger_Path) (models.Passenger_Path, error) {
	err := r.db.Create(&passenger).Error
	if err != nil {
		return models.Passenger_Path{}, err
	}
	return *passenger, nil
}

// Delete implements PassengerRepositoryInt
func (r *passengerRepository) Delete(id int) error {
	err := r.db.Delete(&models.Passenger_Path{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// List implements PassengerRepositoryInt
func (r *passengerRepository) List() ([]models.Passenger_Path, error) {
	passengers := []models.Passenger_Path{}
	var err error

	err = r.db.Find(&passengers).Error

	if err != nil {
		return []models.Passenger_Path{}, err
	}
	return passengers, nil
}

// Read implements PassengerRepositoryInt
func (r *passengerRepository) Read(id int) (models.Passenger_Path, error) {
	passenger := models.Passenger_Path{}
	err := r.db.First(&passenger, id).Error
	if err != nil {
		return models.Passenger_Path{}, err
	}
	return passenger, nil

}

// Update implements PassengerRepositoryInt
func (r *passengerRepository) Update(passenger *models.Passenger_Path) (models.Passenger_Path, error) {
	err := r.db.Save(&passenger).Error
	if err != nil {
		return models.Passenger_Path{}, err
	}
	return *passenger, nil
}
