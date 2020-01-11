package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
)

type VoteGormRepo struct {
	conn *gorm.DB
}

func NewVoteGormRepo(db *gorm.DB) *VoteGormRepo {
	return &VoteGormRepo{conn: db}
}

func (vRepo *VoteGormRepo) CheckVoter(userName string) bool {

	vte := entities.RegParties{}
	errs := vRepo.conn.First(&vte, userName).RecordNotFound()

	if errs == true { //if record not found, return false
		return false
	}

	return true

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
