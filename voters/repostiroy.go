package voters

import "github.com/minas528/Online-voting-System/entities"

type VotersRepository interface {
	Voters() ([]entities.Voters, []error)
	Voter(id uint) (*entities.Voters, []error)
	VoterByGID(gid string) (*entities.Voters,[]error)
	UpdateVoter(voter *entities.Voters) (*entities.Voters, []error)
	Deletevoter(id uint) (*entities.Voters, []error)
	StoreVoter(voter *entities.Voters) (*entities.Voters, []error)
	PhoneExists(uname string) bool
	GIDExists(gid string) bool
	VoterRoles(voters *entities.Voters) ([]entities.Role,[]error)
	Regster4Event(id uint) error
	Vote(pid int,eid int,vid int) []error
	CheckEvent(id uint) (*entities.Events, []error)
	StoreRegVoter(voter *entities.RegVoters) (*entities.RegVoters, []error)
	StoreRegParty(parties *entities.RegParties)(*entities.RegParties,[]error)
	GetRegVoters()([]entities.RegVoters,[]error)
	GetRegParites()([]entities.RegParties,[]error)
	Votes()([]entities.Votes,[]error)
	GetRegVotersByID(id int) (*entities.RegVoters,[]error)
	GetRegPartyByID(id int)(*entities.Parties,[]error)
	GetAlreadyVoted(id int)(*entities.Votes,[]error)

}


type RoleRepository interface {
	Roles() ([]entities.Role, []error)
	Role(id uint) (*entities.Role, []error)
	RoleByName(name string) (*entities.Role, []error)
	UpdateRole(role *entities.Role) (*entities.Role, []error)
	DeleteRole(id uint) (*entities.Role, []error)
	StoreRole(role *entities.Role) (*entities.Role, []error)
}


type SessionRepository interface {
	Session(sessionID string) (*entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}