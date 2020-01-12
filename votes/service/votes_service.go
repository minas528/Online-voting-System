package service

import (
	"github.com/minas528/Online-voting-System/voters"
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/votes"
	"log"
)

type VoterServiceImple struct {
	voterRepo voters.VotersRepository
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
>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db

func NewAuthService(authrepo voters.VotersRepository) *VoterServiceImple {
	return &VoterServiceImple{voterRepo: authrepo}
}
<<<<<<< HEAD
func (asi *VoterServiceImple) Voters() ([]entities.Voters, []error) {
	voters, errs := asi.voterRepo.Voters()
=======

func (vs *VoteServiceImple) IncrementCounter(prtyName string) (*entities.RegParties, []error) { //prtyname comes from button clicked

	vte, errs := vs.voteRepo.GetCounter(prtyName)
	vte.Counter++

	//vte.counter++
	vs.voteRepo.IncrementCounter(vte)

>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db
	if len(errs) > 0 {
		log.Println("serv err")
		return nil, errs
	}
	return voters, errs
}

<<<<<<< HEAD
func (asi *VoterServiceImple) Voter(id uint) (*entities.Voters, []error) {
	vtr, errs := asi.voterRepo.Voter(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return vtr, nil
}
func (asi *VoterServiceImple)VoterByGID(gid string) (*entities.Voters,[]error){
	vtr, errs := asi.voterRepo.VoterByGID(gid)
	if len(errs) >0{
		return nil,errs
	}
	return vtr,errs
}
func (asi *VoterServiceImple) UpdateVoter(pst *entities.Voters) (*entities.Voters, []error) {
	vtr, errs := asi.voterRepo.UpdateVoter(pst)
=======
func (vs *VoteServiceImple) Parties() ([]entities.RegParties, []error) {

	canid, errs := vs.voteRepo.Parties()

>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db
	if len(errs) > 0 {
		return nil, errs
	}
	return vtr, nil
}
func (asi *VoterServiceImple) Deletevoter(id uint) (*entities.Voters, []error) {
	vtr, errs := asi.voterRepo.Deletevoter(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return vtr, nil
}
func (asi *VoterServiceImple) StoreVoter(pst *entities.Voters) (*entities.Voters, []error) {
	vtr, errs := asi.voterRepo.StoreVoter(pst)
	if len(errs) > 0 {
		return nil, errs
	}
	return vtr, nil
}
func (asi *VoterServiceImple)PhoneExists(uname string) bool{
	exists := asi.voterRepo.PhoneExists(uname)
	return exists
}
func (asi *VoterServiceImple)GIDExists(gid string) bool{
	exists := asi.voterRepo.GIDExists(gid)
	return exists
}
func (asi *VoterServiceImple)VoterRoles(voters *entities.Voters) ([]entities.Role,[]error){
	voterRoles, errs := asi.voterRepo.VoterRoles(voters)
	if len(errs) >0{
		return nil,errs
	}
	return voterRoles,errs
}
