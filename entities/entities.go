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
	ID     int
	Name   string `gorm:"type:varchar(255);unique;not null"`
	Logo   string `gorm:"type:varchar(255)"`
	Slogan string `gorm:"type:varchar(255); not null"`
	Event  int
}


type Voters struct {
	ID       int
	FullName string `gorm:"type:varchar(100);unique;not null"`
	GID      string `gorm:"type:varchar(255); not null"` //goverment id
	Password string `gorm:"type:varchar(255); not null"`
	Region   string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	RoleID uint
}
type Votes struct {
	ID uint
	PartiesRefer int
	EventRefer int
	VoterRefer int
}
type Role struct {
	ID uint
	Name string `gorm:"type:varchar(255)"`
	Voter []Voters
}

type RegParties struct {
	ID int
	Parties Parties `gorm:"foreignKey:PartiesRefer"`
	PartiesRefer int
	Event Events `gorm:"foreignkey:EventRefer"`
	EventRefer uint
	Count   int    `gorm:"default:0"`

}

type RegVoters struct {

	ID int
	Voters Voters `gorm:"foreignKey:VoterRefer"`
	VoterRefer int
	flag  bool   `gorm:"default:0"` //did they vote?
	Event Events `gorm:"foreignkey:EventRefer"`
	EventRefer int

}
type Session struct {
	ID uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

type User struct{
	Username string `gorm:"type:varchar(255);unique;not null"`
	ID int
	Region string `gorm:"type:varchar(255)"`
	DID string `gorm:"type:varchar(255)"`
	Age int
	Password string
}
