package votes

import "../../../../github.com/minas528/Online-voting-System/entities"

type VoteService interface {
	CheckVoter(voteID int) bool
	IncrementCounter(prtyName string) (*entities.RegParties, []error)
	Parties() ([]entities.RegParties, []error)
}
