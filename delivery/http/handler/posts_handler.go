package handler

import (
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"../../../../../../github.com/outThabox/Online-voting-System/entities"
	"../../../../../../github.com/outThabox/Online-voting-System/post"
)

type PostsHandler struct {
	tmpl    *template.Template
	pstServ post.PostService
}

func NewPostHandler(T *template.Template, PS post.PostService) *PostsHandler {
	return &PostsHandler{tmpl: T, pstServ: PS}
}

func (ph *PostsHandler) Posts(w http.ResponseWriter, r *http.Request) {
	posts, errs := ph.pstServ.Posts()
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	ph.tmpl.ExecuteTemplate(w, "posts", posts)

}

func (ph *PostsHandler) PostNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pst := entities.Post{}
		pst.Name = r.FormValue("name")
		pst.Writer = r.FormValue("writer")
		pst.Disc = r.FormValue("description")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		pst.Vid = fh.Filename

		CreateFile(&mf, fh.Filename)
		log.Print(pst)

		_, errs := ph.pstServ.StorePost(&pst)
		if errs != nil {
			panic(errs)
		}

		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	} else {
		ph.tmpl.ExecuteTemplate(w, "upload.post", nil)
	}
}

func CreateFile(mf *multipart.File, fname string) {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	path := filepath.Join(wd, "ui", "assets", "vid", fname)

	vid, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer vid.Close()

	io.Copy(vid, *mf)
}
