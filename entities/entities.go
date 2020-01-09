package entities

type Events struct {
<<<<<<< HEAD
	ID          uint
	Name        string `gorm:"type:varchar(255); not null"`
	StartTime   time.Time
	Duration    time.Duration
	EndingTime  time.Time
	Scope       string    `gorm:"type:varchar(255);not null"`
	Competitors []Parties `gorm:"many2many;parties"`
=======
	ID uint
	Name string `gorm:"type:varchar(255); not null"`
	StartTime string
	EndingTime string
	Disc string `gorm:"type:varchar(200);not null"`

>>>>>>> 22a1904d57f4055e35f6cb753c2113699a2fb359
}

type Post struct {
	ID     int
	Name   string `gorm:"type:varchar(200);not null"`
	Writer string `gorm:"type:varchar(200);not null"`
<<<<<<< HEAD
	Disc   string `gorm:"type:varchar(200);not null"`
	Vid    string `gorm:"type:varchar(200);not null"`
=======
	Vid string `gorm:"type:varchar(200);not null"`
	Disc string `gorm:"type:varchar(200);not null"`
>>>>>>> 22a1904d57f4055e35f6cb753c2113699a2fb359
}

type Parties struct {
	Name   string `gorm:"type:varchar(255); not null"`
	Logo   string `gorm:"type:varchar(255)"`
	Slogan string `gorm:"type:varchar(255); not null"`
	Scope  string `gorm:"type:varchar(255); not null"`
}

type Vote struct {
	ID      int  `gorm:"unique;not null"`
	flag    bool `gorm:"default:false"` //did they vote?
	GID     int  //generated id
	PartyID int
}
