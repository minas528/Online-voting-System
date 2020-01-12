package parties

<<<<<<< HEAD
import (
	"github.com/minas528/Online-voting-System/entities"
)
=======
import "../../../../github.com/minas528/Online-voting-System/entities"
>>>>>>> 16e7adbc68177c043a8fc6c3f98223984f6335a7

// PartiesRepository specifies post related database operations
type PartiesRepository interface {
	Parties() ([]entities.Parties, []error)
	Party(id int) (*entities.Parties, []error)
	UpdateParties(Parties *entities.Parties) (*entities.Parties, []error)
	DeleteParties(id int) (*entities.Parties, []error)
	StoreParties(Parties *entities.Parties) (*entities.Parties, []error)
}
<<<<<<< HEAD

type SessionRepository interface {
	Session(sessionID string) (*entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}
=======
>>>>>>> 16e7adbc68177c043a8fc6c3f98223984f6335a7
