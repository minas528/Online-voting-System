package parties
import "github.com/minas528/Online-voting-System/entities"

// PostService specifies post services
type PostService interface {
	Parties() ([]entities.Parties, []error)
	Party(id int) (*entities.Parties, []error)
	UpdateParties(pst *entities.Parties) (*entities.Parties,[]error)
	DeleteParties(id int) (*entities.Parties,[]error)
	StoreParties(pst *entities.Parties) (*entities.Parties,[]error)
}
