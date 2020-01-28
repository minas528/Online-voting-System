package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/Event"
	"github.com/minas528/Online-voting-System/entities"
)

type MockEventRepositoryImple struct {
	Conn *gorm.DB
}

func NewMockEventRepository(conn *gorm.DB) Event.EventRepostory {
	return &MockEventRepositoryImple{Conn: conn}
}

func (meri *MockEventRepositoryImple) Events() ([]entities.Events, []error) {
	events := []entities.Events{entities.EventsMock}
	return events, nil
}

func (meri *MockEventRepositoryImple) Event(id uint) (*entities.Events, []error) {
	event := entities.EventsMock
	return &event, nil
}
func (meri *MockEventRepositoryImple) UpdateEvent(event *entities.Events) (*entities.Events, []error) {
	even := entities.EventsMock
	return &even, nil
}
func (meri *MockEventRepositoryImple) DeleteEvent(id uint) (*entities.Events, []error) {
	even := entities.EventsMock
	if id!=1 {
		return nil, []error{errors.New("Not found")}
	}

	return &even, nil
}
func (meri *MockEventRepositoryImple) StoreEvent(event *entities.Events) (*entities.Events, []error) {
	even := event

	return even, nil
}

