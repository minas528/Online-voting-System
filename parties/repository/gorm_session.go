package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/parties"
)

// SessionGormRepo implements user.SessionRepository interface
type SessionGormRepo struct {
	conn *gorm.DB
}

// NewSessionGormRepo  returns a new SessionGormRepo object
func NewSessionGormRepo(db *gorm.DB) parties.SessionRepository {
	return &SessionGormRepo{conn: db}
}

// Session returns a given stored session
func (sr *SessionGormRepo) Session(sessionID string) (*entities.Session, []error) {
	session := entities.Session{}
	errs := sr.conn.Find(&session, "uuid=?", sessionID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &session, errs
}

// StoreSession stores a given session
func (sr *SessionGormRepo) StoreSession(session *entities.Session) (*entities.Session, []error) {
	sess := session
	errs := sr.conn.Save(sess).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}

// DeleteSession deletes a given session
func (sr *SessionGormRepo) DeleteSession(sessionID string) (*entities.Session, []error) {
	sess, errs := sr.Session(sessionID)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sr.conn.Delete(sess, sess.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}
