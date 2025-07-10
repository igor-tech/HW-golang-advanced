package user

import (
	"errors"
	"order/api/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByPhone(phone string) (*model.User, error) {
	var u model.User
	if err := r.db.Where("phone = ?", phone).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindBySessionId(sid string) (*model.User, error) {
	var u model.User
	if err := r.db.Where("session_id = ?", sid).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpsertSession(u *model.User) error {
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
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error
}
