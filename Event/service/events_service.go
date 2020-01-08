package service

import (
	"github.com/minas528/Online-voting-System/Event"
	"github.com/minas528/Online-voting-System/entities"
	"log"
)

type EventServiceImple struct {
	eventRepo Event.EventRepostory
}

func NewPostService(evnetrepo Event.EventRepostory) *EventServiceImple {
	return &EventServiceImple{eventRepo: evnetrepo}
}
func (esi *EventServiceImple)Posts() ([]entities.Events, []error){
	events,errs := esi.eventRepo.Events()
	if len(errs)>0 {
		log.Println("serv err")
		return nil,errs
	}
	return events,errs
}
func (esi *EventServiceImple) Post(id uint) (*entities.Events, []error){
	evt ,errs := esi.eventRepo.Event(id)
	if len(errs)>0{
		return evt,errs
	}
	return evt,nil
}
func (esi *EventServiceImple)UpdateEvent(events *entities.Events) (*entities.Events,[]error){
	evt,errs:= esi.eventRepo.UpdateEvent(events)
	if len(errs) >0 {
		return nil,errs
	}
	return evt,nil
}
func (esi *EventServiceImple)DeleteEvent(id uint) (*entities.Events,[]error){
	evt,errs := esi.eventRepo.DeleteEvent(id)
	if len(errs)>0 {
		return nil,errs
	}
	return evt,nil
}
func (esi *EventServiceImple)StorePost(events *entities.Events) (*entities.Events,[]error){
	evt,errs := esi.eventRepo.StoreEvent(events)
	if len(errs)>0 {
		return nil,errs
	}
	return evt,nil
}