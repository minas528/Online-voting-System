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