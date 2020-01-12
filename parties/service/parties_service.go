package service

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/parties"
	"log"
)

type PartiesServiceImple struct {
	PartiesRepo parties.PartiesRepository
}

func NewPartiesService(partiesrepo parties.PartiesRepository) *PartiesServiceImple {
	return &PartiesServiceImple{PartiesRepo: partiesrepo}
}
func (ps *PartiesServiceImple) Parties() ([]entities.Parties, []error) {
	Parties, errs := ps.PartiesRepo.Parties()
	if len(errs) > 0 {
		log.Println("serv err")
		return nil, errs
	}
	return Parties, errs
}
func (ps *PartiesServiceImple) Party(id int) (*entities.Parties, []error) {
	pst, errs := ps.PartiesRepo.Party(id)
	if len(errs) > 0 {
		return pst, errs
	}
	return pst, nil
}
func (ps *PartiesServiceImple) UpdateParties(pst *entities.Parties) (*entities.Parties, []error) {
	pst, errs := ps.PartiesRepo.UpdateParties(pst)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, nil
}
func (ps *PartiesServiceImple) DeleteParties(id int) (*entities.Parties, []error) {
	pst, errs := ps.PartiesRepo.DeleteParties(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, nil
}
func (ps *PartiesServiceImple) StoreParties(pst *entities.Parties) (*entities.Parties, []error) {
	pst, errs := ps.PartiesRepo.StoreParties(pst)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, nil
}
