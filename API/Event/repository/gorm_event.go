package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/API/entity"
)

type EventRepositoryImple struct {
	Conn *gorm.DB
}

func NewEventRepository(conn *gorm.DB) *EventRepositoryImple {
	return &EventRepositoryImple{Conn: conn}
}

func (eri *EventRepositoryImple) Events() ([]entity.Events, []error) {
	events := []entity.Events{}
	errs := eri.Conn.Find(&events).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return events, errs
}

func (eri *EventRepositoryImple) Event(id uint) (*entity.Events, []error) {
	event := entity.Events{}
	errs := eri.Conn.First(&event).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &event, errs
}
func (eri *EventRepositoryImple) UpdateEvent(event *entity.Events) (*entity.Events, []error) {
	even := event
	errs := eri.Conn.Save(even).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return even, errs
}
func (eri *EventRepositoryImple) DeleteEvent(id uint) (*entity.Events, []error) {
	even, errs := eri.Event(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = eri.Conn.Delete(even, even.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return even, errs
}
func (eri *EventRepositoryImple) StoreEvent(event *entity.Events) (*entity.Events, []error) {
	even := event
	errs := eri.Conn.Create(even).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return even, errs
}
