package Event

import (
	"github.com/minas528/API/entity"
)

type EventRepostory interface {
	Events() ([]entity.Events, []error)
	Event(id uint) (*entity.Events, []error)
	UpdateEvent(category *entity.Events) (*entity.Events, []error)
	DeleteEvent(id uint) (*entity.Events, []error)
	StoreEvent(category *entity.Events) (*entity.Events, []error)
}
