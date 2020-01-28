package service

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/voters"
	"log"
)

type VoterServiceImple struct {
	voterRepo voters.VotersRepository
}


func NewAuthService(authrepo voters.VotersRepository) *VoterServiceImple {
	return &VoterServiceImple{voterRepo: authrepo}
}
func (asi *VoterServiceImple) Voters() ([]entities.Voters, []error) {
	voters, errs := asi.voterRepo.Voters()
	if len(errs) > 0 {
		log.Println("serv err")
		return nil, errs
	}
	return voters, errs
}
func (asi *VoterServiceImple)CheckEvent(id uint) (*entities.Events, []error){
	events ,errs:= asi.voterRepo.CheckEvent(id)
	if len(errs) >0 {
		return nil,errs

	}
	return events,errs
}

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
func (asi *VoterServiceImple) Regster4Event(id uint) error{
	return nil
}
func (asi *VoterServiceImple) Vote(pid int,eid int,vid int) []error{
	errs := asi.voterRepo.Vote(pid,eid,vid)
	return errs
}
func (asi *VoterServiceImple) StoreRegVoter(voter *entities.RegVoters) (*entities.RegVoters, []error){
	entity ,errs:= asi.voterRepo.StoreRegVoter(voter)
	if len(errs)>0 {
		return nil,errs
	}
	return entity,errs

}

func (asi *VoterServiceImple) StoreRegParty(parties *entities.RegParties)(*entities.RegParties,[]error){
	party ,errs := asi.voterRepo.StoreRegParty(parties)
	if len(errs)>0 {
		return nil,errs
	}
	return party,errs
}
func (asi *VoterServiceImple)GetRegVoters()([]entities.RegVoters,[]error){
	regvoters,errs := asi.voterRepo.GetRegVoters()
	if len(errs) >0{
		return nil,errs
	}
	return regvoters,errs
}
func (asi *VoterServiceImple) GetRegParites()([]entities.RegParties,[]error){
	regparty ,errs:= asi.voterRepo.GetRegParites()
	if len(errs) >0 {
		return nil,errs
	}
	return regparty,errs
}
func (asi *VoterServiceImple) Votes()([]entities.Votes,[]error){
	votes, errs := asi.voterRepo.Votes()
	if len(errs)>0 {
		return nil,errs

	}
	return votes,errs
}

func (asi *VoterServiceImple) GetRegVotersByID(id int) (*entities.RegVoters,[]error){
	regVoter ,errs := asi.voterRepo.GetRegVotersByID(id)
	if len(errs)>0 {
		return nil,errs

	}
	return regVoter,errs
}
func (asi *VoterServiceImple) GetRegPartyByID(id int)(*entities.Parties,[]error){
	regParty,errs:= asi.voterRepo.GetRegPartyByID(id)
	if len(errs)>0 {
		return nil,errs

	}
	return regParty,errs
}

func (asi *VoterServiceImple) GetAlreadyVoted(id int)(*entities.Votes,[]error){
	vtr,errs:= asi.voterRepo.GetAlreadyVoted(id)
	if len(errs) >0 {
		return nil,errs
	}
	return vtr,errs
}