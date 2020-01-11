package service

import (
	"log"

	"../../../../../github.com/minas528/Online-voting-System/entities"
	"../../../../../github.com/minas528/Online-voting-System/post"
)

type UserServiceImple struct {
	userRepo user.UserRepository
}

func NewUserService(userrepo user.UserRepository) *UserServiceImple {
	return &UserServiceImple{userRepo: userrepo}
}
func (us *UserServiceImple) Users() ([]entities.User, []error) {
	users, errs := us.userRepo.Users()
	if len(errs) > 0 {
		log.Println("serv err")
		return nil, errs
	}
	return users, errs
}
func (us *UserServiceImple) User(id int) (*entities.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return usr, errs
	}
	return usr, nil
}
func (us *UserServiceImple) UpdateUser(usr *entities.User) (*entities.User, []error) {
	usr, errs := us.userRepo.UpdatePost(usr)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, nil
}
func (us *UserServiceImple) DeleteUser(id int) (*entities.User, []error) {
	usr, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, nil
}
func (us *UserServiceImple) StoreUser(usr *entities.User) (*entities.User, []error) {
	usr, errs := us.userRepo.StoreUser(usr)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, nil
}
