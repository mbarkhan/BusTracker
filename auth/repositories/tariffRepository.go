package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type TarrifsRepositoryInt interface {
	Create(*models.Tariff) (models.Tariff, error)
	Read(int) (models.Tariff, error)
	Update(*models.Tariff) (models.Tariff, error)
	Delete(int) error
	List() ([]models.Tariff, error)
}

type tarrifsRepository struct {
	db *gorm.DB
}

func NewTarrifsRepository(db *gorm.DB) TarrifsRepositoryInt {
	return &tarrifsRepository{db}
}

func (r *tarrifsRepository) Create(item *models.Tariff) (models.Tariff, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return models.Tariff{}, err
	}
	return *item, nil
}

func (r *tarrifsRepository) Delete(id int) error {
	err := r.db.Delete(&models.Tariff{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *tarrifsRepository) List() ([]models.Tariff, error) {
	items := []models.Tariff{}
	var err error

	err = r.db.Find(&items).Error

	if err != nil {
		return []models.Tariff{}, err
	}
	return items, nil
}

func (r *tarrifsRepository) Read(id int) (models.Tariff, error) {
	item := models.Tariff{}
	err := r.db.First(&item, id).Error
	if err != nil {
		return models.Tariff{}, err
	}
	return item, nil
}

func (r *tarrifsRepository) Update(item *models.Tariff) (models.Tariff, error) {
	err := r.db.Save(&item).Error
	if err != nil {
		return models.Tariff{}, err
	}
	return *item, nil
}
