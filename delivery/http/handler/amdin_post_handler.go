package handler

import (
	"github.com/minas528/Online-voting-System/form"
	"github.com/minas528/Online-voting-System/rtoken"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/post"
)

type AdminPostsHandler struct {
	tmpl    *template.Template
	pstServ post.PostService
	csrfSignKey []byte
}

func NewAdminPostHandler(T *template.Template, PS post.PostService,csKey []byte) *AdminPostsHandler {
	return &AdminPostsHandler{tmpl: T, pstServ: PS,csrfSignKey:csKey}
}

func (ph *AdminPostsHandler) Posts(w http.ResponseWriter, r *http.Request) {
	posts, errs := ph.pstServ.Posts()
	if len(errs) >0 {
		log.Println(errs)
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(ph.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	templDate := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Posts []entities.Post
		CSRF  string
	}{
		Values:nil,
		VErrors:nil,
		Posts:posts,
		CSRF:token,
	}
	ph.tmpl.ExecuteTemplate(w, "admin.posts", templDate)

}

func (ph *AdminPostsHandler) PostNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ph.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet{
		newPostForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF string
		}{
			Values:nil,
			VErrors:nil,
			CSRF:token,
		}
		ph.tmpl.ExecuteTemplate(w,"admin.upload.post",newPostForm)
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		newPostForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		newPostForm.Required("name", "writer")
		newPostForm.MinLength("description", 20)
		newPostForm.CSRF = token

		if !newPostForm.Valid() {
			ph.tmpl.ExecuteTemplate(w, "admin.upload.post", newPostForm)
			return
		}

		mf, fh, err := r.FormFile("vid")
		if err != nil {
			newPostForm.VErrors.Add("vid", "File error")
			ph.tmpl.ExecuteTemplate(w, "admin.upload.post", newPostForm)
			return
		}

		defer mf.Close()

		pst := &entities.Post{
			Name:   r.FormValue("name"),
			Writer: r.FormValue("writer"),
			Disc:   r.FormValue("description"),
			Vid:    fh.Filename,
		}
		writeFile(&mf, fh.Filename)

		_, errs := ph.pstServ.StorePost(pst)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)

	}
}

func (ph *AdminPostsHandler) AdminPostsUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ph.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idraw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idraw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		pst, errs := ph.pstServ.Post(id)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		values := url.Values{}
		values.Add("id",idraw)
		values.Add("name",pst.Name)
		values.Add("description",pst.Disc)
		values.Add("writer",pst.Writer)
		values.Add("vid",pst.Vid)

		upPstForm := struct {
			Values   url.Values
			VErrors  form.ValidationErrors
			Post *entities.Post
			CSRF     string
		}{
			Values:   values,
			VErrors:  form.ValidationErrors{},
			Post: pst,
			CSRF:     token,
		}
		ph.tmpl.ExecuteTemplate(w, "upload.post", upPstForm)
		return
	}
	//TODO
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		updatePostForm := form.Input{Values:r.PostForm,VErrors:form.ValidationErrors{}}
		updatePostForm.Required("name","description")
		updatePostForm.MinLength("description",20)
		updatePostForm.CSRF = token

		pstId , err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w,http.StatusText(http.StatusBadRequest),http.StatusBadRequest)
		}



		pst := &entities.Post{
			ID: int(pstId),
			Name:r.FormValue("name"),
			Writer:r.FormValue("writer"),
			Disc:r.FormValue("description"),
			Vid:r.FormValue("vid"),
		}
		mf, fh, err := r.FormFile("vid")
		if err == nil {
			pst.Vid = fh.Filename
			err = writeFile(&mf, pst.Vid)
		}
		if mf != nil {
			defer mf.Close()
		}
		_, errs := ph.pstServ.UpdatePost(pst)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
		return

	}
}

func (ph *AdminPostsHandler) AdminPostDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil{
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	_, errs := ph.pstServ.DeletePost(int(id))
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "ui", "assets", "vid", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
