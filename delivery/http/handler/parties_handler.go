package handler

import (
	"github.com/minas528/Online-voting-System/parties"
	"html/template"
	"log"
	"net/http"
)

type PartiesHandler struct {
	tmpl    *template.Template
	pstServ parties.PartiesService
}

func NewPartiesHandler(T *template.Template, PS parties.PartiesService) *AdminPartiesHandler {
	return &AdminPartiesHandler{tmpl: T, pstServ: PS}
}

func (ph *AdminPartiesHandler) Parties(w http.ResponseWriter, r *http.Request) {
	party, errs := ph.pstServ.Parties()

	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	log.Println(party)
	ph.tmpl.ExecuteTemplate(w, "parties", party)

}
