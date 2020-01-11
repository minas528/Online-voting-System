package handler

import (
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"../../../../../../github.com/minas528/Online-voting-System/entities"
	"../../../../../../github.com/minas528/Online-voting-System/authentication"
)

type UsersHandler struct {
	tmpl    *template.Template
	usrServ user.UserService
}

func NewUserHandler(T *template.Template, US user.UserService) *UsersHandler {
	return &UsersHandler{tmpl: T, usrServ: PS}
}

func (uh *UsersHandler) Users(w http.ResponseWriter, r *http.Request) {
	users, errs := uh.usrServ.Users()
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	uh.tmpl.ExecuteTemplate(w, "signup", nil)

}

func (uh *UsersHandler) UserNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		usr := entities.User{}
		usr.Name = r.FormValue("name")
		usr.Writer = r.FormValue("writer")
		usr.Disc = r.FormValue("description")

		mf, fh, err := r.FormFile("vid")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		usr.Vid = fh.Filename

		CreateFile(&mf, fh.Filename)
		log.Print(usr)

		_, errs := uh.usrServ.StoreUser(&usr)
		if errs != nil {
			panic(errs)-
		}

		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		ph.tmpl.ExecuteTemplate(w, "upload.post", nil)
	}
}

/*func CreateFile(mf *multipart.File, fname string) {
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
*/
