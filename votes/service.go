package votes

import "../../../../github.com/minas528/Online-voting-System/entities"

type VoteService interface {
	CheckVoter(userName string) bool
	IncrementCounter(prtyName string) (*entities.RegParties, []error)
}
