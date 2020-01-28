package entity

type Events struct {
	ID         uint    `json:"id"`
	Name       string `gorm:"type:varchar(255); not null" json:"name"`
	StartTime  string  `json:"start_time"`
	EndingTime string	`json:"ending_time"`
	Disc       string `gorm:"type:varchar(200);not null" "json:"disc""`
}



 