package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type CompanyUserRepositoryInt interface {
	Create(*models.Company_Users) (models.Company_Users, error)
	Read(int) (models.Company_Users, error)
	Update(*models.Company_Users) (models.Company_Users, error)
	Delete(int) error
	List() ([]models.Company_Users, error)
}

type companyUserRepository struct {
	db *gorm.DB
}

func NewCompanyUserRepository(db *gorm.DB) CompanyUserRepositoryInt {
	return &companyUserRepository{db}
}
func (r *companyUserRepository) Create(company_user *models.Company_Users) (models.Company_Users, error) {
	err := r.db.Create(&company_user).Error
	if err != nil {
		return models.Company_Users{}, err
	}
	return *company_user, nil
}

func (r *companyUserRepository) Delete(id int) error {
	err := r.db.Delete(&models.Company_Users{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *companyUserRepository) List() ([]models.Company_Users, error) {
	comp_users := []models.Company_Users{}
	var err error

	err = r.db.Find(&comp_users).Error

	if err != nil {
		return []models.Company_Users{}, err
	}
	return comp_users, nil
}

func (r *companyUserRepository) Read(id int) (models.Company_Users, error) {
	comp_user := models.Company_Users{}
	err := r.db.First(&comp_user, id).Error
	if err != nil {
		return models.Company_Users{}, err
	}
	return comp_user, nil
}

func (r *companyUserRepository) Update(comp_user *models.Company_Users) (models.Company_Users, error) {
	err := r.db.Save(&comp_user).Error
	if err != nil {
		return models.Company_Users{}, err
	}
	return *comp_user, nil
}
