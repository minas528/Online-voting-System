package authentication

import "../../../../github.com/minas528/Online-voting-System/entities"

type UserService interface {
	Users() ([]entities.User, []error)
	User(id int) (*entities.User, []error)
	UpdateUser(user *entities.User) (*entities.User, []error)
	DeleteUser(id int) (*entities.User, []error)
	StoreUser(user *entities.User) (*entities.User, []error)
}
