package service

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/post"
	"log"


)

type PostServiceImple struct {
	postRepo post.PostRepository
}

func NewPostService(postrepo post.PostRepository) *PostServiceImple {
	return &PostServiceImple{postRepo: postrepo}
}
func (ps *PostServiceImple) Posts() ([]entities.Post, []error) {
	posts, errs := ps.postRepo.Posts()
	if len(errs) > 0 {
		log.Println("serv err")
		return nil, errs
	}
	return posts, errs
}
func (ps *PostServiceImple) Post(id int) (*entities.Post, []error) {
	pst, errs := ps.postRepo.Post(id)
	if len(errs) > 0 {
		return pst, errs
	}
	return pst, nil
}
func (ps *PostServiceImple) UpdatePost(pst *entities.Post) (*entities.Post, []error) {
	pst, errs := ps.postRepo.UpdatePost(pst)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, nil
}
func (ps *PostServiceImple) DeletePost(id int) (*entities.Post, []error) {
	pst, errs := ps.postRepo.DeletePost(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, nil
}
func (ps *PostServiceImple) StorePost(pst *entities.Post) (*entities.Post, []error) {
	pst, errs := ps.postRepo.StorePost(pst)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, nil
}
