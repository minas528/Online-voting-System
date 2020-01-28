package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/minas528/API/Event"
	"github.com/minas528/API/entity"
	"net/http"
	"strconv"
)

type EventHandler struct {
	eveserv Event.EventService
}

func NewEventHandler( ES Event.EventService) *EventHandler {
	return &EventHandler{ eveserv: ES}
}

func (ph *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	events, errs := ph.eveserv.Events()
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err := json.MarshalIndent(events,"","\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (ph *EventHandler) GetSingleEvent(w http.ResponseWriter, r *http.Request, pr httprouter.Params){

	id, err := strconv.Atoi(pr.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	event,errs:=ph.eveserv.Event(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(event, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (ph *EventHandler) PostEvent(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	event := &entity.Events{}

	err := json.Unmarshal(body, event)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	events,errs := ph.eveserv.StoreEvent(event)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/events/%d", events.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

func (ph *EventHandler) PutEvent(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	id, err := strconv.Atoi(pr.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	event,errs:= ph.eveserv.Event(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &event)

	event ,errs = ph.eveserv.UpdateEvent(event)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(event, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (eh *EventHandler) DeleteEvents(w http.ResponseWriter, r *http.Request,pr httprouter.Params) {

	id, err := strconv.Atoi(pr.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_,errs := eh.eveserv.DeleteEvent(uint(id))
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return

}

