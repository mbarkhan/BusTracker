package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type UserCreditRepositoryInt interface {
	Create(*models.UserCredit) (models.UserCredit, error)
	Read(uint) ([]models.UserCredit, error)
}

type userCreditRepository struct {
	db *gorm.DB
}

func NewUserCreditRepository(db *gorm.DB) UserCreditRepositoryInt {
	return &userCreditRepository{db}
}

func (r *userCreditRepository) Create(userCredit *models.UserCredit) (models.UserCredit, error) {
	err := r.db.Create(&userCredit).Error
	if err != nil {
		return models.UserCredit{}, err
	}
	return *userCredit, nil
}
func (r *userCreditRepository) Read(userid uint) ([]models.UserCredit, error) {
	userCredits := []models.UserCredit{}
	err := r.db.Model(&models.UserCredit{}).Where("f_user_id=?", userid).Find(&userCredits).Error
	if err != nil {
		return []models.UserCredit{}, err
	}
	return userCredits, nil
}
