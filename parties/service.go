package parties

import "github.com/minas528/Online-voting-System/entities"

// PostService specifies post services
type PostService interface {
	Parties() ([]entities.Parties, []error)
	Party(id int) (*entities.Parties, []error)
	UpdateParties(pst *entities.Parties) (*entities.Parties, []error)
	DeleteParties(id int) (*entities.Parties, []error)
	StoreParties(pst *entities.Parties) (*entities.Parties, []error)
}

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionID string) (*entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}
