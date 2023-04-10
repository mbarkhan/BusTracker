//TO DO LIST
/*
* Update detaledbookID while creating user by accounting service (CREATE)
*(DELETE) Control deleting user-> not used in any transactions
*
 */

package services

import (
	utils "githubmbarkhanBusTracker/auth/Utils"
	"githubmbarkhanBusTracker/auth/models"
	"githubmbarkhanBusTracker/auth/repositories"
)

type UserServiceInt interface {
	Create(*models.User) (models.User, error)
	Read(int) (models.User, error)
	Update(*models.User) (models.User, error)
	Delete(int) error
	List() ([]models.User, error)
	GetUserByMobile(string) (*models.User, error)
	UserSubmittedCompanies(int) (*[]models.Company_Users, error)
}

type userService struct {
	user repositories.UserRepositoryInt
}

func NewUserService(iuser repositories.UserRepositoryInt) UserServiceInt {
	return &userService{user: iuser}
}

func (r *userService) Create(user *models.User) (models.User, error) {
	var err error
	//Password Encryption
	user.Password, err = utils.EncryptPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	return r.user.Create(user)
}

func (r *userService) Read(userId int) (models.User, error) {
	return r.user.Read(userId)
}

func (r *userService) Update(user *models.User) (models.User, error) {
	return r.user.Update(user)
}

func (r *userService) Delete(userId int) error {
	return r.user.Delete(userId)
}

func (r *userService) List() ([]models.User, error) {
	return r.user.List()
}

func (r *userService) GetUserByMobile(mobNum string) (*models.User, error) {
	return r.user.GetUserByMobile(mobNum)
}

// UserSubmittedCompanies implements UserServiceInt
func (r *userService) UserSubmittedCompanies(userID int) (*[]models.Company_Users, error) {
	return r.user.UserSubmittedCompanies(userID)
}
