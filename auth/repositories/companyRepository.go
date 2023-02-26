package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type CompanyRepositoryInt interface {
	Create(*models.Company) (models.Company, error)
	Read(int) (models.Company, error)
	Update(*models.Company) (models.Company, error)
	Delete(int) error
	List() ([]models.Company, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepositoryInt {
	return &companyRepository{db}
}

func (r *companyRepository) Create(company *models.Company) (models.Company, error) {
	err := r.db.Create(&company).Error
	if err != nil {
		return models.Company{}, err
	}
	return *company, nil
}
func (r *companyRepository) Read(id int) (models.Company, error) {
	company := models.Company{}
	err := r.db.First(&company, id).Error
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (r *companyRepository) Update(company *models.Company) (models.Company, error) {
	err := r.db.Save(&company).Error
	if err != nil {
		return models.Company{}, err
	}
	return *company, nil
}

func (r *companyRepository) Delete(id int) error {
	err := r.db.Delete(&models.Company{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *companyRepository) List() ([]models.Company, error) {
	company := []models.Company{}
	var err error

	err = r.db.Find(&company).Error

	if err != nil {
		return []models.Company{}, err
	}
	return company, nil
}
