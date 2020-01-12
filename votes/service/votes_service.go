package service

import (
	"github.com/minas528/Online-voting-System/voters"
	"github.com/minas528/Online-voting-System/entities"
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
