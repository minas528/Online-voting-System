package Event

import (
	"github.com/minas528/Online-voting-System/entities"
)

type EventRepostory interface {
	Events() ([]entities.Events, []error)
	Event(id uint) (*entities.Events, []error)
	UpdateEvent(category *entities.Events) (*entities.Events, []error)
	DeleteEvent(id uint) (*entities.Events, []error)
	StoreEvent(category *entities.Events) (*entities.Events, []error)
}

