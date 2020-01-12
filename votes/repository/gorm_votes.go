package repository

import (
<<<<<<< HEAD
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
	"log"
=======
<<<<<<< HEAD
=======
	"log"

	"../../../../../github.com/minas528/Online-voting-System/entities"
>>>>>>> 16e7adbc68177c043a8fc6c3f98223984f6335a7
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db
)

type VoterGormRepo struct {
	conn *gorm.DB
}

func NewVoterGormRepo(conn *gorm.DB) *VoterGormRepo  {
	return &VoterGormRepo{conn:conn}
}

<<<<<<< HEAD
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
	
=======
func (vRepo *VoteGormRepo) CheckVoter(voteID int) (*entities.RegVoters, []error) {

	vte := entities.RegVoters{}
	errs := vRepo.conn.First(&vte, voteID).GetErrors()

	/*	if errs == true { //if record not found, return false
		return false
	}*/
	if len(errs) > 0 {
		log.Println("faliled fetching this voter")
		return nil, errs
	}

	return &vte, errs

>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db
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
	log.Println("now here")
	return votee, errs
}

<<<<<<< HEAD
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
=======
func (vRepo *VoteGormRepo) Parties() ([]entities.RegParties, []error) {
	canids := []entities.RegParties{}
	errs := vRepo.conn.Find(&canids).GetErrors()
>>>>>>> 90ea9b8aaea637f705c6fe5b924c293b64b367db

	if len(errs) >0 {
		return nil,errs
	}
	return voterRoles,errs
}
