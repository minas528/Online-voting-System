package post

import "../../../../github.com/minas528/Online-voting-System/entities"

// PostRepository specifies post related database operations
type PostRepository interface {
	Posts() ([]entities.Post, []error)
	Post(id int) (*entities.Post, []error)
	UpdatePost(post *entities.Post) (*entities.Post, []error)
	DeletePost(id int) (*entities.Post, []error)
	StorePost(post *entities.Post) (*entities.Post, []error)
}
