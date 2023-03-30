package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type UserCreditRepositoryInt interface {
	Create(*models.User_Credit) (models.User_Credit, error)
	Read(int) ([]models.User_Credit, error)
}

type userCreditRepository struct {
	db *gorm.DB
}

func NewUserCreditRepository(db *gorm.DB) UserCreditRepositoryInt {
	return &userCreditRepository{db}
}

func (r *userCreditRepository) Create(userCredit *models.User_Credit) (models.User_Credit, error) {
	err := r.db.Create(&userCredit).Error
	if err != nil {
		return models.User_Credit{}, err
	}
	return *userCredit, nil
}
func (r *userCreditRepository) Read(userid int) ([]models.User_Credit, error) {
	userCredits := []models.User_Credit{}
	err := r.db.Model(&models.User_Credit{}).Where("f_user_id=?", userid).Find(&userCredits).Error
	if err != nil {
		return []models.User_Credit{}, err
	}
	return userCredits, nil
}
