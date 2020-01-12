package repository

import (
	"log"

	"../../../../../github.com/minas528/Online-voting-System/entities"

	"github.com/jinzhu/gorm"
)

type VoteGormRepo struct {
	conn *gorm.DB
}

func NewVoteGormRepo(db *gorm.DB) *VoteGormRepo {
	return &VoteGormRepo{conn: db}
}

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

}

func (vRepo *VoteGormRepo) GetCounter(prtyName string) (*entities.RegParties, []error) {
	vte := entities.RegParties{}
	errs := vRepo.conn.First(&vte, prtyName).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &vte, errs
}
func (vRepo *VoteGormRepo) IncrementCounter(vote *entities.RegParties) (*entities.RegParties, []error) {
	vte := vote
	errs := vRepo.conn.Save(vte).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return vte, errs
}

func (vRepo *VoteGormRepo) Parties() ([]entities.RegParties, []error) {
	canids := []entities.RegParties{}
	errs := vRepo.conn.Find(&canids).GetErrors()

	if len(errs) > 0 {
		log.Println("faliled from gorm_vote")
		return nil, errs
	}
	return canids, errs
}
