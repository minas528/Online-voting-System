package repository

import (
	"log"

	"../../../../../github.com/minas528/Online-voting-System/entities"

	"github.com/jinzhu/gorm"
)

type UserGormRepo struct {
	conn *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) *UserGormRepo {
	return &UserGormRepo{conn: db}
}

func (uRepo *UserGormRepo) Posts() ([]entities.User, []error) {
	users := []entities.User{}
	errs := uRepo.conn.Find(&user).GetErrors()
	log.Println("in side of users repo")

	if len(errs) > 0 {
		log.Println("failed")
		return nil, errs
	}
	return users, errs
}

func (uRepo *UserGormRepo) User(id int) (*entities.User, []error) {
	usr := entities.User{}
	errs := uRepo.conn.First(&usr, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &usr, errs
}
func (uRepo *UserGormRepo) UpdateUser(user *entities.User) (*entities.User, []error) {
	usr := user
	errs := uRepo.conn.Save(usr).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (uRepo *User) Deleteuser(id int) (*entities.User, []error) {
	usr, errs := uRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (uRepo *UserGormRepo) StoreUser(post *entities.User) (*entities.User, []error) {
	usr := user
	
	errs := uRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
