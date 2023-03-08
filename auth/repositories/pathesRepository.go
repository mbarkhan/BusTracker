package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type PathesRepositoryInt interface {
	Create(*models.Pathes) (models.Pathes, error)
	Read(int) (models.Pathes, error)
	Update(*models.Pathes) (models.Pathes, error)
	Delete(int) error
	List() ([]models.Pathes, error)
}

type pathesRepository struct {
	db *gorm.DB
}

func NewPathesRepository(db *gorm.DB) PathesRepositoryInt {
	return &pathesRepository{db}
}

func (r *pathesRepository) Create(path *models.Pathes) (models.Pathes, error) {
	err := r.db.Create(&path).Error
	if err != nil {
		return models.Pathes{}, err
	}
	return *path, nil
}

func (r *pathesRepository) Delete(id int) error {
	err := r.db.Delete(&models.Pathes{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *pathesRepository) List() ([]models.Pathes, error) {
	pathes := []models.Pathes{}
	var err error

	err = r.db.Find(&pathes).Error

	if err != nil {
		return []models.Pathes{}, err
	}
	return pathes, nil
}

func (r *pathesRepository) Read(id int) (models.Pathes, error) {
	path := models.Pathes{}
	err := r.db.First(&path, id).Error
	if err != nil {
		return models.Pathes{}, err
	}
	return path, nil
}

func (r *pathesRepository) Update(path *models.Pathes) (models.Pathes, error) {
	err := r.db.Save(&path).Error
	if err != nil {
		return models.Pathes{}, err
	}
	return *path, nil
}
