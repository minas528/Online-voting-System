package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
)

type MockSessionGormRepo struct {
	conn *gorm.DB
}


func NewMockSessionGormRepo(db *gorm.DB) *MockSessionGormRepo {
	return &MockSessionGormRepo{conn: db}
}
// Session returns a given stored session
func (msgr *MockRoleGormRepo) Session(sessionID string) (*entities.Session, []error) {
	session := entities.SessionMock
	return &session, nil
}

// StoreSession stores a given session
func (msgr *MockRoleGormRepo) StoreSession(session *entities.Session) (*entities.Session, []error) {
	sess := session

	return sess, nil
}

// DeleteSession deletes a given session
func (msgr *MockRoleGormRepo) DeleteSession(sessionID string) (*entities.Session, []error) {
	sess := entities.SessionMock
	if sessionID != "1"{
		return nil, []error{errors.New("Not found")}
	}
	return &sess, nil
}

