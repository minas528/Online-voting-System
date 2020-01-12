package post

import (
	"../../../../github.com/minas528/Online-voting-System/entities"
)

// PostService specifies post services
type PostService interface {
	Posts() ([]entities.Post, []error)
	Post(id int) (*entities.Post, []error)
	UpdatePost(pst *entities.Post) (*entities.Post, []error)
	DeletePost(id int) (*entities.Post, []error)
	StorePost(pst *entities.Post) (*entities.Post, []error)
}
