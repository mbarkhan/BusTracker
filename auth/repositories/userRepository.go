package repositories

import (
	"githubmbarkhanBusTracker/auth/models"

	"gorm.io/gorm"
)

type UserRepositoryInt interface {
	Create(*models.User) (models.User, error)
	Read(int) (models.User, error)
	Update(*models.User) (models.User, error)
	Delete(int) error
	List() ([]models.User, error)
	GetUserByMobile(string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInt {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}
func (r *userRepository) Read(id int) (models.User, error) {
	user := models.User{}
	err := r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (r *userRepository) Delete(id int) error {
	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List() ([]models.User, error) {
	users := []models.User{}
	var err error

	err = r.db.Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}
func (r *userRepository) GetUserByMobile(mobNum string) (models.User, error) {

	user := models.User{}
	error := r.db.Model(models.User{}).Where("mobile=?", mobNum).First(&user).Error

	if error != nil {
		return models.User{}, error
	}

	return user, nil
}
