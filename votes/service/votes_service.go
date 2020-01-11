package service

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/votes"
)

type VoteServiceImple struct {
	voteRepo votes.VoteRepository
}

func NewVoteService(voterepo votes.VoteRepository) *VoteServiceImple {
	return &VoteServiceImple{voteRepo: voterepo}
}

func (vs *VoteServiceImple) CheckVoter(voteID int) bool {

	voter, _ := vs.voteRepo.CheckVoter(voteID)
	/*if check{
		vs.IncrementCounter
	}*/
	//if check is true, run increment counter on this page, if false, display voter
	if voter.Flag == 0 { //if voter hasnt voted, return false
		return false
	}
	return true

}

func (vs *VoteServiceImple) IncrementCounter(prtyName string) (*entities.RegParties, []error) { //prtyname comes from button clicked

	vte, errs := vs.voteRepo.GetCounter(prtyName)
	vte.Counter++

	//vte.counter++
	vs.voteRepo.IncrementCounter(vte)

	if len(errs) > 0 {
		return nil, errs
	}
	return vte, nil
}

func (vs *VoteServiceImple) Parties() ([]entities.RegParties, []error) {

	canid, errs := vs.voteRepo.Parties()

	if len(errs) > 0 {

		return nil, errs
	}
	return canid, errs
}
