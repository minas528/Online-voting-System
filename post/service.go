package post

<<<<<<< HEAD
import "github.com/minas528/Online-voting-System/entities"
=======
import (
	"../../../../github.com/minas528/Online-voting-System/entities"
)
>>>>>>> 345e47e2cd443abfe01044e928281286fec9a418

// PostService specifies post services
type PostService interface {
	Posts() ([]entities.Post, []error)
	Post(id int) (*entities.Post, []error)
	UpdatePost(pst *entities.Post) (*entities.Post, []error)
	DeletePost(id int) (*entities.Post, []error)
	StorePost(pst *entities.Post) (*entities.Post, []error)
}
