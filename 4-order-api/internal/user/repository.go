package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByPhone(phone string) (*User, error) {
	var u User
	if err := r.db.Where("phone = ?", phone).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindBySessionId(sid string) (*User, error) {
	var u User
	if err := r.db.Where("session_id = ?", sid).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpsertSession(u *User) error {
	existing, err := r.FindByPhone(u.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.db.Create(u).Error
		}
		return err
	}

	existing.SessionID = u.SessionID
	existing.Code = u.Code
	return r.db.Save(existing).Error
}

func (r *UserRepository) UpdateFields(id uint,
	fields map[string]interface{}) error {

	return r.db.
		Model(&User{}).
		Where("id = ?", id).
		Updates(fields).Error
}
