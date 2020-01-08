package repository

import "github.com/jinzhu/gorm"

type EventRepository struct {
	Conn *gorm.DB
}

func NewEventRepository(conn *gorm.DB) *EventRepository  {
	return &EventRepository{Conn:conn}
}


