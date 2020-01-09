package post

import "../../../../github.com/outThabox/Online-voting-System/entities"

// PostService specifies post services
type PostService interface {
	Posts() ([]entities.Post, []error)
	Post(id int) (*entities.Post, []error)
<<<<<<< HEAD
	UpdatePost(category *entities.Post) (*entities.Post, []error)
	DeletePost(id int) (*entities.Post, []error)
	StorePost(category *entities.Post) (*entities.Post, []error)
=======
	UpdatePost(pst *entities.Post) (*entities.Post,[]error)
	DeletePost(id int) (*entities.Post,[]error)
	StorePost(pst *entities.Post) (*entities.Post,[]error)
>>>>>>> 22a1904d57f4055e35f6cb753c2113699a2fb359
}
