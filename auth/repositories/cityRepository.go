package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type CityRepositoryInt interface {
	Create(*models.City) (models.City, error)
	Read(int) (models.City, error)
	Update(*models.City) (models.City, error)
	Delete(int) error
	List() ([]models.City, error)
}

type citiesRepository struct {
	db *gorm.DB
}

func NewCitiesRepository(db *gorm.DB) CityRepositoryInt {
	return &citiesRepository{db}
}

func (r *citiesRepository) Create(city *models.City) (models.City, error) {
	err := r.db.Create(&city).Error
	if err != nil {
		return models.City{}, err
	}
	return *city, nil
}

func (r *citiesRepository) Delete(id int) error {
	err := r.db.Delete(&models.City{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *citiesRepository) List() ([]models.City, error) {
	cities := []models.City{}
	var err error

	err = r.db.Find(&cities).Error

	if err != nil {
		return []models.City{}, err
	}
	return cities, nil
}

func (r *citiesRepository) Read(id int) (models.City, error) {
	city := models.City{}
	err := r.db.First(&city, id).Error
	if err != nil {
		return models.City{}, err
	}
	return city, nil
}

func (r *citiesRepository) Update(city *models.City) (models.City, error) {
	err := r.db.Save(&city).Error
	if err != nil {
		return models.City{}, err
	}
	return *city, nil
}
