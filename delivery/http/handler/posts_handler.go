package handler

import (
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/minas528/Online-voting-System/post"
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

<<<<<<< HEAD
/*func (ph *PostsHandler) PostNew(w http.ResponseWriter, r *http.Request) {
=======
func (ph *PostsHandler) choseParty(w http.ResponseWriter, r *http.Request) {
>>>>>>> 16e7adbc68177c043a8fc6c3f98223984f6335a7
	if r.Method == http.MethodPost {
		pst := entities.Post{}
		pst.Name = r.FormValue("name")
		pst.Writer = r.FormValue("writer")
		pst.Disc = r.FormValue("description")

		mf, fh, err := r.FormFile("vid")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		pst.Vid = fh.Filename

		CreateFiles(&mf, fh.Filename)
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
*/
func CreateFiles(mf *multipart.File, fname string) {
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
