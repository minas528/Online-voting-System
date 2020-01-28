package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/web_project/Online-voting-System/post"
	"github.com/pkg/errors"
)

type MockPostGormRepo struct {
	conn *gorm.DB
}

func NewMockPostGormRepo(db *gorm.DB) post.PostRepository {
	return &MockPostGormRepo{conn: db}
}

func (pRepo *MockPostGormRepo) Posts() ([]entities.Post, []error) {
	posts := []entities.Post{entities.PostMock}

	return posts, nil
}

func (pRepo *MockPostGormRepo) Post(id int) (*entities.Post, []error) {
	pst := entities.PostMock
	return &pst, nil
}
func (pRepo *MockPostGormRepo) UpdatePost(post *entities.Post) (*entities.Post, []error) {
	pst := entities.PostMock
	return &pst, nil
}
func (pRepo *MockPostGormRepo) DeletePost(id int) (*entities.Post, []error) {
	pst := entities.PostMock
	if id != 1{
		return nil,[]error{errors.New("Not Found")}
	}
	return &pst, nil
}

func (pRepo *MockPostGormRepo) StorePost(post *entities.Post) (*entities.Post, []error) {
	pst := post

	return pst, nil
}

