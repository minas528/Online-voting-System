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
	ID int
	Name   string `gorm:"type:varchar(255);unique;not null"`
	Logo   string `gorm:"type:varchar(255)"`
	Slogan string `gorm:"type:varchar(255); not null"`
	Scope  string `gorm:"type:varchar(255); not null"`
}

type Voters struct {
	ID       uint
	FullName string `gorm:"type:varchar(100);unique;not null"`
	GID      string `gorm:"type:varchar(255); not null"` //goverment id
	Password string `gorm:"type:varchar(255); not null"`
	Region   string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	RoleID uint
	Votes []Votes
}
type Votes struct {
	ID uint
	VID string
	UserID uint
	PartiesID uint
}
type Role struct {
	ID uint
	Name string `gorm:"type:varchar(255)"`
	Voter []Voters
}

type RegParties struct {
	ID int
	Logo      string `gorm:"type:varchar(255)"`
	Motto     string `gorm:"type:varchar(255); not null"`
	PartyName string `gorm:"type:varchar(255); not null"`
	Counter   int    `gorm:"default:0"`
	Event Events `gorm:"foreignkey:EventRefer"`
	EventRefer int
}

type RegVoters struct {
	ID int
	Uname string `gorm:"type:varchar(100);unique"`
	flag  bool   `gorm:"default:0"` //did they vote?
	VotingID string `gorm:"type:varchar(20);unique"`
	Event Events `gorm:"foreignkey:EventRefer"`
	EventRefer int

}
type Session struct {
	ID uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

