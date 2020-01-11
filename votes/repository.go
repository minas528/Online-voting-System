package votes

import "../../../../github.com/minas528/Online-voting-System/entities"

type VoteRepository interface {
	CheckVoter(voteID int) (*entities.RegVoters, []error)
	IncrementCounter(vte *entities.RegParties) (*entities.RegParties, []error)
	GetCounter(prtyName string) (*entities.RegParties, []error)
	Parties() ([]entities.RegParties, []error)
}
