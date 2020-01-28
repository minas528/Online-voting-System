package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/parties"
)




type MockPartiesGormRepo struct {
	conn *gorm.DB
}
func NewMockPartiesGormRepo(conn *gorm.DB) parties.PartiesRepository  {
	return &MockPartiesGormRepo{conn:conn}
}

func (mockPartiesRepo *MockPartiesGormRepo) Parties() ([]entities.Parties, []error) {
	parties := []entities.Parties{entities.PartiesMock}
	return parties,nil
}

func (mockPartiesRepo *MockPartiesGormRepo) Party(id int) (*entities.Parties, []error) {
	party := entities.PartiesMock
	return &party,nil
}

func (mockPartiesRepo *MockPartiesGormRepo) UpdateParties(Parties *entities.Parties) (*entities.Parties, []error) {
	p := entities.PartiesMock
	return &p,nil
}

func (mockPartiesRepo *MockPartiesGormRepo) DeleteParties(id int) (*entities.Parties, []error) {
	r := entities.PartiesMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &r, nil
}

func (mockPartiesRepo *MockPartiesGormRepo) StoreParties(Parties *entities.Parties) (*entities.Parties, []error) {
	r := Parties
	return r,nil
}


