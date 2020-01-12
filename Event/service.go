package Event

import "github.com/minas528/Online-voting-System/entities"

type EventService interface {
	Events() ([]entities.Events, []error)
	Event(id uint) (*entities.Events, []error)
	UpdateEvent(event *entities.Events) (*entities.Events, []error)
	DeleteEvent(id uint) (*entities.Events, []error)
	StoreEvent(event *entities.Events) (*entities.Events, []error)
}
