package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
	"log"
)

type VoterGormRepo struct {
	conn *gorm.DB
}

func NewVoterGormRepo(conn *gorm.DB) *VoterGormRepo  {
	return &VoterGormRepo{conn:conn}
}

func (ari *VoterGormRepo)Voters() ([]entities.Voters, []error){
	voters := []entities.Voters{}
	errs := ari.conn.Find(&voters).GetErrors()
	if len(errs)> 0{
		return nil,errs
	}
	return voters,errs
	
}
func (ari *VoterGormRepo)Voter(id uint) (*entities.Voters, []error){
	voter := entities.Voters{}
	errs := ari.conn.First(&voter).GetErrors()
	if len(errs) >0{
		return nil,errs
	}
	return &voter,errs
	
}

func (ari *VoterGormRepo) VoterByGID(gid string) (*entities.Voters,[]error){
	voter := entities.Voters{}
	errs := ari.conn.Find(&voter,"g_id=?",gid).GetErrors()
	if len(errs) >0 {
		return nil,errs
	}
	return &voter,errs
}
func (ari *VoterGormRepo)UpdateVoter(voter *entities.Voters) (*entities.Voters, []error){
	vote := voter
	errs := ari.conn.Save(vote).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return vote, errs	
}
func (ari *VoterGormRepo)Deletevoter(id uint) (*entities.Voters, []error){
	voter ,errs := ari.Voter(id)
	errs = ari.conn.Delete(voter).GetErrors()
	if len(errs)>0{
		return nil,errs
	}
	return voter,errs
}
func (ari *VoterGormRepo)StoreVoter(voter *entities.Voters) (*entities.Voters, []error){
	votee := voter

	errs := ari.conn.Create(votee).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	log.Println("admin addes")
	return votee, errs
}

func (ari *VoterGormRepo) PhoneExists(phone string) bool{
	voter := entities.Voters{}
	errs := ari.conn.Find(&voter,"phone=?",phone).GetErrors()
	if len(errs) >0 {
		return false
	}
	return true
}
func (ari *VoterGormRepo)GIDExists(gid string) bool{
	voter := entities.Voters{}
	errs := ari.conn.Find(&voter,"g_id=?",gid).GetErrors()
	if len(errs) >0 {
		return false
	}
	return true
}
func (ari *VoterGormRepo)VoterRoles(voters *entities.Voters) ([]entities.Role,[]error){
	voterRoles := []entities.Role{}
	errs := ari.conn.Model(voters).Related(&voterRoles).GetErrors()

	if len(errs) >0 {
		return nil,errs
	}
	return voterRoles,errs
}
func (ari *VoterGormRepo) Regster4Event(id uint) error{
	return nil
}
func (ari *VoterGormRepo) Vote(pid int,eid int,vid int) []error{
	c_voter,_ := ari.GetAlreadyVoted(vid)

	if c_voter == nil{
		vote := entities.Votes{
			VoterRefer:vid,
			PartiesRefer:pid,
			EventRefer:eid,
		}

		errs := ari.conn.Create(&vote).GetErrors()
		if len(errs) >0{
			return errs
		}
	}
	return nil
}
func (ari *VoterGormRepo) CheckEvent(id uint) (*entities.Events, []error){
	events := entities.Events{}
	errs := ari.conn.Find(&events,"id=?",id).GetErrors()
	if len(errs) >0{
		return nil, errs
	}
	return &events,errs
}
func (ari *VoterGormRepo) StoreRegVoter(voter *entities.RegVoters) (*entities.RegVoters, []error){

	events:= voter
	errs:= ari.conn.Save(events).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}

	return events,errs
}

func (ari *VoterGormRepo) StoreRegParty(parties *entities.RegParties)(*entities.RegParties,[]error){
	party := parties
	errs := ari.conn.Save(parties).GetErrors()
	if len(errs) > 0 {
		return nil,errs
	}
	return party ,errs
}
func (ari *VoterGormRepo) GetRegVoters()([]entities.RegVoters,[]error){
	regVoter := []entities.RegVoters{}
	errs := ari.conn.Find(regVoter).GetErrors()
	if len(errs)>0{
		return nil,errs
	}
	return regVoter,errs
}
func(ari *VoterGormRepo) GetRegParites()([]entities.RegParties,[]error){
	regParty := []entities.RegParties{}
	errs := ari.conn.Find(regParty).GetErrors()
	if len(errs)>0{
		return nil,errs
	}
	return regParty,errs
}
func (ari *VoterGormRepo) Votes()([]entities.Votes,[]error){
	votes := []entities.Votes{}
	errs := ari.conn.Find(votes).GetErrors()
	if len(errs) >0 {
		return nil,errs
	}
	return votes,errs
}

func (ari *VoterGormRepo) GetRegVotersByID(id int) (*entities.RegVoters,[]error){
	votee := entities.RegVoters{}
	errs := ari.conn.Find(&votee,"voter_refer=?",id).GetErrors()
	if len(errs) >0 {
		return nil,errs
	}
	return &votee,errs
}
func (ari *VoterGormRepo) GetRegPartyByID(id int)(*entities.Parties,[]error){
	party := entities.Parties{}
	errs := ari.conn.Find(&party,"id=?",id).GetErrors()
	if len(errs) >0 {
		return nil,errs
	}
	return &party,errs
}

func (ari *VoterGormRepo) GetAlreadyVoted(id int)(*entities.Votes,[]error){
	voted := entities.Votes{}
	errs := ari.conn.Find(&voted,"voter_refer=?",id).GetErrors()
	if len(errs) >0 {
		return nil,errs
	}
	return &voted,errs

}