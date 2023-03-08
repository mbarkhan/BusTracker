//TO DO LIST
/*
* Update detaledbookID while creating user by accounting service (CREATE)
*(DELETE) Control deleting user-> not used in any transactions
*
 */

package services

import "githubmbarkhanBusTracker/auth/models"

type UserServiceInt interface {
	Create(*models.User) (models.User, error)
	Read(int) (models.User, error)
	Update(*models.User) (models.User, error)
	Delete(int) error
	List() ([]models.User, error)
}

type userService struct {
	user UserServiceInt
}

func NewUserService(iuser UserServiceInt) UserServiceInt {
	return &userService{user: iuser}
}

func (r *userService) Create(user *models.User) (models.User, error) {
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
