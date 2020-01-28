package service

import (
	"github.com/minas528/API/Event"
	"github.com/minas528/API/entity"
	"log"

)

type EventServiceImple struct {
	eventRepo Event.EventRepostory
}

func NewEventService(evnetrepo Event.EventRepostory) *EventServiceImple {
	return &EventServiceImple{eventRepo: evnetrepo}
}
func (esi *EventServiceImple) Events() ([]entity.Events, []error) {
	events, errs := esi.eventRepo.Events()
	if len(errs) > 0 {
		log.Println("serv err")
		return nil, errs
	}
	return events, errs
}
func (esi *EventServiceImple) Event(id uint) (*entity.Events, []error) {
	evt, errs := esi.eventRepo.Event(id)
	if len(errs) > 0 {
		return evt, errs
	}
	return evt, nil
}
func (esi *EventServiceImple) UpdateEvent(events *entity.Events) (*entity.Events, []error) {
	evt, errs := esi.eventRepo.UpdateEvent(events)
	if len(errs) > 0 {
		return nil, errs
	}
	return evt, nil
}
func (esi *EventServiceImple) DeleteEvent(id uint) (*entity.Events, []error) {
	evt, errs := esi.eventRepo.DeleteEvent(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return evt, nil
}
func (esi *EventServiceImple) StoreEvent(events *entity.Events) (*entity.Events, []error) {
	evt, errs := esi.eventRepo.StoreEvent(events)
	if len(errs) > 0 {
		return nil, errs
	}
	return evt, nil
}
