package Event

import "github.com/minas528/API/entity"

type EventService interface {
	Events() ([]entity.Events, []error)
	Event(id uint) (*entity.Events, []error)
	UpdateEvent(event *entity.Events) (*entity.Events, []error)
	DeleteEvent(id uint) (*entity.Events, []error)
	StoreEvent(event *entity.Events) (*entity.Events, []error)
}
