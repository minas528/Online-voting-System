package repository

import (
"github.com/minas528/Online-voting-System/entities"
"log"

"github.com/jinzhu/gorm"
)

type PartiesGormRepo struct {
	conn *gorm.DB
}

func NewPostGormRepo(db *gorm.DB) *PartiesGormRepo {
	return &PartiesGormRepo{conn: db}
}

func (pRepo *PartiesGormRepo) Parties() ([]entities.Parties, []error) {
	party := []entities.Parties{}
	errs := pRepo.conn.Find(&party).GetErrors()
	log.Println("in side of posts repo")

	if len(errs) > 0 {
		log.Println("faliled")
		return nil, errs
	}
	return party, errs
}

func (pRepo *PartiesGormRepo) Party(id int) (*entities.Parties, []error) {
	pst := entities.Parties{}
	errs := pRepo.conn.First(&pst, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &pst, errs
}
func (pRepo *PartiesGormRepo) UpdateParties(Parties *entities.Parties) (*entities.Parties, []error) {
	pst := Parties
	errs := pRepo.conn.Save(pst).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
func (pRepo *PartiesGormRepo) DeleteParties(id int) (*entities.Parties, []error) {
	pst, errs := pRepo.Party(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
func (pRepo *PartiesGormRepo) StoreParties(Parties *entities.Parties) (*entities.Parties, []error) {
	pst := Parties

	errs := pRepo.conn.Create(pst).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
