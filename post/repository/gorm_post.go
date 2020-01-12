package repository

import (
	"log"

	"github.com/minas528/Online-voting-System/entities"

	"github.com/jinzhu/gorm"
)

type PostGormRepo struct {
	conn *gorm.DB
}

func NewPostGormRepo(db *gorm.DB) *PostGormRepo {
	return &PostGormRepo{conn: db}
}

func (pRepo *PostGormRepo) Posts() ([]entities.Post, []error) {
	posts := []entities.Post{}
	errs := pRepo.conn.Find(&posts).GetErrors()
	log.Println("in side of posts repo")

	if len(errs) > 0 {
		log.Println("failed")
		return nil, errs
	}
	return posts, errs
}
 
func (pRepo *PostGormRepo) Post(id int) (*entities.Post, []error) {
	pst := entities.Post{}
	errs := pRepo.conn.First(&pst, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &pst, errs
}
func (pRepo *PostGormRepo) UpdatePost(post *entities.Post) (*entities.Post, []error) {
	pst := post
	errs := pRepo.conn.Save(pst).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
func (pRepo *PostGormRepo) DeletePost(id int) (*entities.Post, []error) {
	pst, errs := pRepo.Post(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

func (pRepo *PostGormRepo) StorePost(post *entities.Post) (*entities.Post, []error) {
	pst := post

	errs := pRepo.conn.Create(pst).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
