package entities

import (
	"time"
)

type Events struct {
	ID uint
	Name string `gorm:"type:varchar(255); not null"`
	StartTime time.Time
	Duration time.Duration
	EndingTime time.Time
	Scope string `gorm:"type:varchar(255);not null"`
	Competitors []Parties `gorm:"many2many;parties"`
}

type Post struct {
	Name string `gorm:"type:varchar(200);not null"`
	Writer string `gorm:"type:varchar(200);not null"`
	Disc string `gorm:"type:varchar(200);not null"`
	Vid string `gorm:"type:varchar(200);not null"`
}

type Parties struct {
	Name string `gorm:"type:varchar(255); not null"`
	Logo string `gorm:"type:varchar(255)"`
	Slogan string `gorm:"type:varchar(255); not null"`
	Scope string `gorm:"type:varchar(255); not null"`
}
