package handler

import (
	"html/template"
	"log"
	"net/http"

	"../../../../../../github.com/minas528/Online-voting-System/entities"
	
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
		usr.Name = r.FormValue("username")
		usr.Writer = r.FormValue("DID")
		usr.Disc = r.FormValue("region")

		_, errs := uh.usrServ.StoreUser(&usr)
		if errs != nil {
			panic(errs)
		}

		http.Redirect(w, r, "/signup", http.StatusSeeOther)
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
