package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (r *sessionsRepo) AddSessions(session model.Session) error {
	err := r.db.Create(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionsRepo) DeleteSession(token string) error {
	err := r.db.Delete(&model.Session{}, "token = ?", token).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionsRepo) UpdateSessions(session model.Session) error {
	result := r.db.Model(&model.Session{}).
		Where("email = ?", session.Email).
		Update("token", session.Token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	var session model.Session
	err := r.db.First(&session, "email = ?", email).Error
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (r *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := r.db.First(&session, "token = ?", token).Error
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (r *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
