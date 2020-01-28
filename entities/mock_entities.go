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
	Event:  1,
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
	EventRefer: 1,
	PartiesRefer: 1,
	VoterRefer: 1,
}
var RoleMock = Role {
	ID: 1,
	Name: "",
}

var RegPartiesMock = RegParties {
	ID: 1,
	PartiesRefer :  1,
	Parties :  Parties{},
	EventRefer: 1,
	Count:   1,
	Event: Events{},
}

var RegVotersMock =RegVoters {
	ID: 1,
	VoterRefer: 1,
	Voters:Voters{},
	flag:  false,
	Event: Events{},
	EventRefer: 1,

}
var SessionMock = Session {
	ID: 1,
	UUID : "",
	Expires :   1,
	SigningKey: []byte{},
}


