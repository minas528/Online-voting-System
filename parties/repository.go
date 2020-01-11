package parties

import "../../../../github.com/minas528/Online-voting-System/entities"

// PostRepository specifies post related database operations
type PartiesRepository interface {
	Parties() ([]entities.Parties, []error)
	Party(id int) (*entities.Parties, []error)
	UpdateParties(Parties *entities.Parties) (*entities.Parties, []error)
	DeleteParties(id int) (*entities.Parties, []error)
	StoreParties(Parties *entities.Parties) (*entities.Parties, []error)
}
