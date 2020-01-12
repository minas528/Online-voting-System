package handler

import (
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"../../../../../../github.com/minas528/Online-voting-System/entities"
	"../../../../../../github.com/minas528/Online-voting-System/post"
)

type AdminPostsHandler struct {
	tmpl    *template.Template
	pstServ post.PostService
}

func NewAdminPostHandler(T *template.Template, PS post.PostService) *AdminPostsHandler {
	return &AdminPostsHandler{tmpl: T, pstServ: PS}
}

func (ph *AdminPostsHandler) Posts(w http.ResponseWriter, r *http.Request) {
	posts, errs := ph.pstServ.Posts()
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	ph.tmpl.ExecuteTemplate(w, "admin.posts", posts)

}

func (ph *AdminPostsHandler) PostNew(w http.ResponseWriter, r *http.Request) {
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

func (ph *AdminPostsHandler) AdminPostsUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idraw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idraw)
		if err != nil {
			panic(err)
		}
		pst, errs := ph.pstServ.Post(id)
		if len(errs) > 0 {
			panic(errs)
		}
		ph.tmpl.ExecuteTemplate(w, "admin.posts.update", pst)
	} else if r.Method == http.MethodPost {
		pst := entities.Post{}
		pst.ID, _ = strconv.Atoi(r.FormValue("id"))
		pst.Name = r.FormValue("name")
		pst.Writer = r.FormValue("writer")
		pst.Disc = r.FormValue("description")
		pst.Vid = r.FormValue("vid")

		mf, _, err := r.FormFile("vid")

		if err != nil {
			panic(err)
		}
		defer mf.Close()
		writeFile(&mf, pst.Vid)

		_, errs := ph.pstServ.UpdatePost(&pst)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
	}
}

func (ph *AdminPostsHandler) AdminPostDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		_, errs := ph.pstServ.DeletePost(id)
		if len(errs) > 0 {
			panic(errs)
		}

	}
	http.Redirect(w, r, "admin/posts", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "vid", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
