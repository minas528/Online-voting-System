package handler

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/parties"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type AdminPartiesHandler struct {
	tmpl    *template.Template
	pstServ parties.PartiesService
}

func NewAdminPartiesHandler(T *template.Template, PS parties.PartiesService) *AdminPartiesHandler {
	return &AdminPartiesHandler{tmpl: T, pstServ: PS}
}

func (ph *AdminPartiesHandler) Parties(w http.ResponseWriter, r *http.Request) {
	party, errs := ph.pstServ.Parties()
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	ph.tmpl.ExecuteTemplate(w, "parties", party)

}

func (ph *AdminPartiesHandler) PartiesNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pst := entities.Parties{}
		pst.Name = r.FormValue("name")
		pst.Slogan = r.FormValue("slogan")
		//pst.Scope = r.FormValue("description")

		mf, fh, err := r.FormFile("vid")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		pst.Logo = fh.Filename

		CreateFile(&mf, fh.Filename)
		log.Print(pst)

		_, errs := ph.pstServ.StoreParties(&pst)
		if errs != nil {
			panic(errs)
		}

		http.Redirect(w, r, "/parties", http.StatusSeeOther)
	} else {
		ph.tmpl.ExecuteTemplate(w, "upload.party", nil)
	}
}
/*
func (ph *AdminPartiesHandler) AdminPartiesUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idraw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idraw)
		if err != nil {
			panic(err)
		}
		pst, errs := ph.pstServ.Party(id)
		if len(errs) > 0 {
			panic(errs)
		}
		ph.tmpl.ExecuteTemplate(w, "admin.parties.update.html", pst)
	} else if r.Method == http.MethodPost {
		pst := entities.Parties{}
		pst.ID, _ = strconv.Atoi(r.FormValue("id"))
		pst.Name = r.FormValue("catname")
		pst.Slogan = r.FormValue("writer")
		//pst.scope = r.FormValue("description")
		pst.Logo = r.FormValue("vid")

		mf, _, err := r.FormFile("vid")

		if err != nil {
			panic(err)
		}
		defer mf.Close()
		writeFiles(&mf, pst.Logo)

		_, errs := ph.pstServ.UpdateParties(&pst)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w, r, "/admin/parties", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/parties", http.StatusSeeOther)
	}
}

func (ph *AdminPartiesHandler) AdminPartiesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		_, errs := ph.pstServ.DeleteParties(id)
		if len(errs) > 0 {
			panic(errs)
		}

	}
	http.Redirect(w, r, "admin/parties", http.StatusSeeOther)
}
*/

/*
func writeFiles(mf *multipart.File, fname string) {

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
*/


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
