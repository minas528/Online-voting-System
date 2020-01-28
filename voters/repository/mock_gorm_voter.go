package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
)

type MockGormVoter struct {
	conn *gorm.DB
}

func (mgv *MockGormVoter) StoreRegParty(parties *entities.RegParties) (*entities.RegParties, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) GetRegVoters() ([]entities.RegVoters, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) GetRegParites() ([]entities.RegParties, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) GetRegVotersByID(id int) (*entities.RegVoters, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) GetRegPartyByID(id int) (*entities.Parties, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) GetAlreadyVoted(id int) (*entities.Votes, []error) {
	panic("implement me")
}

func NewMockGormvoter(conn *gorm.DB) *MockGormVoter  {
	return &MockGormVoter{conn:conn}

}

func (mgv *MockGormVoter) Voters() ([]entities.Voters, []error){
	voters :=[]entities.Voters{entities.VotersMock}
	return voters,nil

}
func (mgv *MockGormVoter) Voter(id uint) (*entities.Voters, []error){
	voter := entities.VotersMock
	return &voter,nil

}

func (mgv *MockGormVoter) VoterByGID(gid string) (*entities.Voters,[]error){
	voter := entities.VotersMock
	return &voter,nil
}
func (mgv *MockGormVoter) UpdateVoter(voter *entities.Voters) (*entities.Voters, []error){
	vote := entities.VotersMock

	return &vote, nil
}
func (mgv *MockGormVoter) Deletevoter(id uint) (*entities.Voters, []error){
	voter  := entities.VotersMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &voter,nil
}
func (mgv *MockGormVoter) StoreVoter(voter *entities.Voters) (*entities.Voters, []error){
	votee := voter
	return votee, nil
}

func (mgv *MockGormVoter) PhoneExists(phone string) bool{
	voter := entities.VotersMock
	if voter.Phone != phone{
		return false
	}
	return true
}
func (mgv *MockGormVoter) GIDExists(gid string) bool{
	voter := entities.VotersMock
	if voter.GID != gid {
		return false
	}
	return true
}
func (mgv *MockGormVoter) VoterRoles(voters *entities.Voters) ([]entities.Role,[]error){
	voterRoles := []entities.Role{entities.RoleMock}
	return voterRoles,nil
}
func (mgv *MockGormVoter) Regster4Event(id uint) error{
	return nil
}
func (mgv *MockGormVoter) Vote(pid int,eid int,vid int) []error{
	return nil
}
func (mgv *MockGormVoter) CheckEvent(id uint) (*entities.Events, []error){
	events := entities.EventsMock
	return &events,nil
}
func (mgv *MockGormVoter) StoreRegVoter(voter *entities.RegVoters) (*entities.RegVoters, []error){
	events:= voter
	return events,nil
}
func (mgv *MockGormVoter) Votes() ([]entities.Votes, []error) {
	panic("implement me")
}


func (mgv *MockGormVoter) VoteByParty(pName string) ([]entities.Votes, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) CheckVoter(voteID int) (*entities.RegVoters, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) IncrementCounter(vte *entities.RegParties) (*entities.RegParties, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) GetCounter(prtyName string) (*entities.RegParties, []error) {
	panic("implement me")
}

func (mgv *MockGormVoter) Parties() ([]entities.RegParties, []error) {
	panic("implement me")
}

