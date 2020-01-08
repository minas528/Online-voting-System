package Event

import "github.com/minas528/Online-voting-System/entities"

type EventRepostory interface {
	Events() ([]entities.Events,error)
	Event(id int) (entities.Events,error)
}
