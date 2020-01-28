package entities


var EventsMock = Events {
	ID:         1,
	Name :      "Mock Event Name",
	StartTime:  "Mock Event StartTime",
	EndingTime:  "Mock Event EndingTime",
	Disc:      "Mock Event Disc",
}
var PostMock = Post {
	ID:     1,
	Name:   "Mock Post Name",
	Writer: "Mock post writer",
	Vid:  "Mock vid",
	Disc: "Mock post description",
}

var PartiesMock = Parties {
	ID: 1,
	Name : "Mock party name",
	Logo :  "Mock party logo",
	Slogan: "Mock party slogan",
	Scope:  "Mock party scope",
}

var VotersMock = Voters {
	ID  :     1,
	FullName: "",
	GID :     "",
	Password: "",
	Region:   "",
	Phone :   "",
	RoleID: 1,
}
var VotesMock = Votes {
	ID: 1,
	VID: "",
	UserID: 1,
	PartiesID: 1,
}
var RoleMock = Role {
	ID: 1,
	Name: "",
}

var RegPartiesMock = RegParties {
	ID: 1,
	Logo :  "",
	Motto :  "",
	PartyName: "",
	Counter:   1,
	Event: Events{},
	EventRefer: 1,
}

var RegVotersMock =RegVoters {
	ID: 1,
	Uname: "",
	flag:  false,
	VotingID: "",
	Event: Events{},
	EventRefer: 1,

}
var SessionMock = Session {
	ID: 1,
	UUID : "",
	Expires :   1,
	SigningKey: []byte{},
}


