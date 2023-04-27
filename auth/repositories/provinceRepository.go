package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type ProvinceRepositoryInt interface {
	Create(*models.Province) (models.Province, error)
	Read(int) (models.Province, error)
	Update(*models.Province) (models.Province, error)
	Delete(int) error
	List() ([]models.Province, error)
}

type provincesRepository struct {
	db *gorm.DB
}

func NewProvincesRepository(db *gorm.DB) ProvinceRepositoryInt {
	return &provincesRepository{db}
}

func (r *provincesRepository) Create(province *models.Province) (models.Province, error) {
	err := r.db.Create(&province).Error
	if err != nil {
		return models.Province{}, err
	}
	return *province, nil
}

func (r *provincesRepository) Delete(id int) error {
	err := r.db.Delete(&models.Province{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *provincesRepository) List() ([]models.Province, error) {
	provinces := []models.Province{}
	var err error

	err = r.db.Find(&provinces).Error

	if err != nil {
		return []models.Province{}, err
	}
	return provinces, nil
}

func (r *provincesRepository) Read(id int) (models.Province, error) {
	province := models.Province{}
	err := r.db.First(&province, id).Error
	if err != nil {
		return models.Province{}, err
	}
	return province, nil
}

func (r *provincesRepository) Update(province *models.Province) (models.Province, error) {
	err := r.db.Save(&province).Error
	if err != nil {
		return models.Province{}, err
	}
	return *province, nil
}
