package votes


import (
	"github.com/minas528/Online-voting-System/entities"
)

type VotersService interface {
	Voters() ([]entities.Voters, []error)
	Voter(id uint) (*entities.Voters, []error)
	VoterByGID(gid string) (*entities.Voters, []error)
	UpdateVoter(voter *entities.Voters) (*entities.Voters, []error)
	Deletevoter(id uint) (*entities.Voters, []error)
	StoreVoter(voter *entities.Voters) (*entities.Voters, []error)
	PhoneExists(uname string) bool
	GIDExists(gid string) bool
	VoterRoles(voters *entities.Voters) ([]entities.Role, []error)
	StoreRegVoter(voter *entities.Voters) (*entities.Voters, []error)
}

type VoteService interface {
	CheckVoter(voteID int) bool
	IncrementCounter(prtyName string) (*entities.RegParties, []error)
	Parties() ([]entities.RegParties, []error)
}


type RoleService interface {
	Roles() ([]entities.Role, []error)
	Role(id uint) (*entities.Role, []error)
	RoleByName(name string) (*entities.Role, []error)
	UpdateRole(role *entities.Role) (*entities.Role, []error)
	DeleteRole(id uint) (*entities.Role, []error)
	StoreRole(role *entities.Role) (*entities.Role, []error)
}


type SessionService interface {
	Session(sessionID string) (*entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}