package voters

<<<<<<< HEAD
import "../../../../github.com/minas528/Online-voting-System/entities"
=======
<<<<<<< HEAD
import (
	"github.com/minas528/Online-voting-System/entities"
)

type VotersService interface {
	Voters() ([]entities.Voters, []error)
	Voter(id uint) (*entities.Voters, []error)
	VoterByGID(gid string) (*entities.Voters,[]error)
	UpdateVoter(voter *entities.Voters) (*entities.Voters, []error)
	Deletevoter(id uint) (*entities.Voters, []error)
	StoreVoter(voter *entities.Voters) (*entities.Voters, []error)
	PhoneExists(uname string) bool
	GIDExists(gid string) bool
	VoterRoles(voters *entities.Voters) ([]entities.Role,[]error)
=======
import "github.com/minas528/Online-voting-System/entities"
>>>>>>> 8ba8eb050e4b504cae99e995a2fe7e64222d1378

type VoteService interface {
	CheckVoter(voteID int) bool
	IncrementCounter(prtyName string) (*entities.RegParties, []error)
	Parties() ([]entities.RegParties, []error)
>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db
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