package service

import (
	"../../../../../github.com/minas528/Online-voting-System/entities"
	"../../../../../github.com/minas528/Online-voting-System/votes"
)

type VoteServiceImple struct {
	voteRepo votes.VoteRepository
}

func (vs *VoteServiceImple) CheckVoter(userName string) bool {

	check := vs.voteRepo.CheckVoter(userName)
	//if check is true, run increment counter on this page, if false, display voter
	return check

}

func (vs *VoteServiceImple) IncrementCounter(prtyName string) (*entities.RegParties, []error) { //prtyname comes from button clicked

	vte, errs := vs.voteRepo.GetCounter(prtyName)
	vte.Counter++

	vs.CheckVoter("")
	//vte.counter++
	vs.voteRepo.IncrementCounter(vte)

	if len(errs) > 0 {
		return nil, errs
	}
	return vte, nil
}

func (vs *VoteServiceImple) Canidates() ([]entities.RegParties, []error) {

	canid, errs := vs.voteRepo.Canidates()

	if len(errs) > 0 {

		return nil, errs
	}
	return canid, errs
}
