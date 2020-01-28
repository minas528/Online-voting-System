package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
)

type VoteGormRepo struct {
	conn *gorm.DB
}

func NewVoteGormRepo(conn *gorm.DB) *VoteGormRepo  {
	return &VoteGormRepo{conn:conn}
}


func (vgr *VoteGormRepo) Votes() ([]entities.Votes, []error){
	votes := []entities.Votes{}
	errs := vgr.conn.Find(&votes).GetErrors()
	if len(errs)>0{
		return nil, errs
	}
	return votes, nil
}
func (vgr *VoteGormRepo) Vote(id uint) (*entities.Votes, []error){
	vote := entities.Votes{}
	errs := vgr.conn.First(&vote).GetErrors()
	if len(errs) >0{
		return nil,errs
	}
	return &vote,errs
}
func(vgr *VoteGormRepo) VoteByParty(id int) ([]entities.Votes,[]error){
	votes := []entities.Votes{}
	errs:= vgr.conn.Find(&votes,"parties_id",id).GetErrors()
	if len(errs)>0{
		return nil,errs
	}
	return votes,errs
}
func (vgr *VoteGormRepo) CheckVoter(voteID int) (*entities.RegVoters, []error){

}
func (vgr *VoteGormRepo) IncrementCounter(vte *entities.RegParties) (*entities.RegParties, []error){

}
func (vgr *VoteGormRepo) GetCounter(prtyName string) (*entities.RegParties, []error){

}
func (vgr *VoteGormRepo) Parties() ([]entities.RegParties, []error){

}

