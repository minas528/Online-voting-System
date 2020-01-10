package entities

type Events struct {
	ID         uint
	Name       string `gorm:"type:varchar(255); not null"`
	StartTime  string
	EndingTime string
	Disc       string `gorm:"type:varchar(200);not null"`
}
type Post struct {
	ID     int
	Name   string `gorm:"type:varchar(200);not null"`
	Writer string `gorm:"type:varchar(200);not null"`

	Vid  string `gorm:"type:varchar(200);not null"`
	Disc string `gorm:"type:varchar(200);not null"`
}

type Parties struct {
	Name   string `gorm:"type:varchar(255);unique;not null"`
	Logo   string `gorm:"type:varchar(255)"`
	Slogan string `gorm:"type:varchar(255); not null"`
	Scope  string `gorm:"type:varchar(255); not null"`
}

type Voters struct {
	Uname    string `gorm:"type:varchar(100);unique;not null"`
	GID      string `gorm:"type:varchar(255); not null"` //goverment id
	Password string `gorm:"type:varchar(255); not null"`
}

type RegParties struct {
	logo      string `gorm:"type:varchar(255)"`
	motto     string `gorm:"type:varchar(255); not null"`
	partyName string `gorm:"type:varchar(255); not null"`
	counter   int    `gorm:"default:0"`
}

type RegVoters struct {
	Uname string `gorm:"type:varchar(100);unique"`
	flag  bool   `gorm:"default:0"` //did they vote?
}
