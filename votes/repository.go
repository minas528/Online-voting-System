package votes

import "github.com/minas528/Online-voting-System/entities"


type VoteRepository interface {
	Votes() ([]entities.Votes, []error)
	Vote(id uint) (*entities.Votes, []error)
	VoteByParty(pName string) ([]entities.Votes,[]error)
	CheckVoter(voteID int) (*entities.RegVoters, []error)
	IncrementCounter(vte *entities.RegParties) (*entities.RegParties, []error)
	GetCounter(prtyName string) (*entities.RegParties, []error)
	Parties() ([]entities.RegParties, []error)
}
