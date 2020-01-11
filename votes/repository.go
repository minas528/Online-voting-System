package votes

import "github.com/minas528/Online-voting-System/entities"

type VoteRepository interface {
	CheckVoter(userName string) (check bool)
	IncrementCounter(vte *entities.RegParties) (*entities.RegParties, []error)
	GetCounter(prtyName string) (*entities.RegParties, []error)
}
