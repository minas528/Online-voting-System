package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/minas528/Online-voting-System/Event"
	"github.com/minas528/Online-voting-System/entities"
)

type EventHandler struct {
	tmpl    *template.Template
	eveserv Event.EventService
}

func NewEventHandler(T *template.Template, ES Event.EventService) *EventHandler {
	return &EventHandler{tmpl: T, eveserv: ES}
}

func (ph *EventHandler) Events(w http.ResponseWriter, r *http.Request) {
	events, errs := ph.eveserv.Events()
	print(len(events))
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}
	ph.tmpl.ExecuteTemplate(w, "events", events)

}

func (ph *EventHandler) EventNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		evt := entities.Events{}
		evt.Name = r.FormValue("name")
		evt.StartTime = r.FormValue("start_date")
		evt.EndingTime = r.FormValue("end_date")

		_, errs := ph.eveserv.StoreEvent(&evt)
		if errs != nil {
			panic(errs)
		}

		http.Redirect(w, r, "/events", http.StatusSeeOther)
	} else {
		ph.tmpl.ExecuteTemplate(w, "new.event", nil)
	}
}

func (eh *EventHandler) UpdateEvents(w http.ResponseWriter, r *http.Request) {

}

func (eh *EventHandler) DeleteEvents(w http.ResponseWriter, r *http.Request) {

}
