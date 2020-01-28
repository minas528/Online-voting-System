package handler

import (
	"github.com/minas528/web_project/Online-voting-System/votes"
	"html/template"
	"log"
	"net/http"

	_ "github.com/minas528/Online-voting-System/votes/service"
)

type VotesHandler struct {
	tmpl    *template.Template
	vteServ votes.VoteService
}

func NewVotesHandler(T *template.Template, VS votes.VoteService) *VotesHandler {
	return &VotesHandler{tmpl: T, vteServ: VS}
}

func (vh *VotesHandler) Vote(w http.ResponseWriter, r *http.Request) {
	canis, errs := vh.vteServ.Parties()
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	vh.tmpl.ExecuteTemplate(w, "vote.form", canis)

}
func (vh *VotesHandler) Chose(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//vte := entities.RegParties{}
		//vte.PartyName = r.FormValue("partyname")
		//
		//_, errs := vh.vteServ.IncrementCounter(vte.PartyName)
		//
		//if errs != nil {
		//	panic(errs)
		//}
		//
		////http.Redirect(w, r, "/vote", http.StatusSeeOther)
	}

}

