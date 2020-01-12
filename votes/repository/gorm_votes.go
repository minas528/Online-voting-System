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
