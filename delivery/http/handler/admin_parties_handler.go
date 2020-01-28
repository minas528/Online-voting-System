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

	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/parties"
)

type AdminPartiesHandler struct {
	tmpl    *template.Template
	pstServ parties.PartiesService
}

func NewAdminPartiesHandler(T *template.Template, PS parties.PartiesService) *AdminPartiesHandler {
	return &AdminPartiesHandler{tmpl: T, pstServ: PS}
}

func (ph *AdminPartiesHandler) AdminParties(w http.ResponseWriter, r *http.Request) {
	party, errs := ph.pstServ.Parties()

	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	log.Println(party)
	ph.tmpl.ExecuteTemplate(w, "admin.parties.layout", party)

}

func (ph *AdminPartiesHandler) PartiesNew(w http.ResponseWriter, r *http.Request) {
	//log.Println("why is this")
	if r.Method == http.MethodPost {
		prt := entities.Parties{}
		prt.Name = r.FormValue("name")
		prt.Slogan = r.FormValue("slogan")
		prt.Event = 1
		//pst.Scope = r.FormValue("description")

		mf, fh, err := r.FormFile("vid")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		prt.Logo = fh.Filename

		//log.Println("logging 2")

		CreateFile(&mf, fh.Filename)
		//log.Print(prt)

		_, errs := ph.pstServ.StoreParties(&prt)
		if errs != nil {
			panic(errs)
		}

		//log.Println("logging")
		http.Redirect(w, r, "/party", http.StatusSeeOther)
	} else {
		ph.tmpl.ExecuteTemplate(w, "upload.party", nil)
	}
}


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
		CreateFile(&mf, pst.Logo)

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



func CreateFile(mf *multipart.File, fname string) {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	path := filepath.Join(wd, "ui", "assets", "party", fname)

	vid, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer vid.Close()

	io.Copy(vid, *mf)
}
