package store

import (
	"errors"
	"go-echo-starter/model"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) All() ([]model.User, error) {
	var u []model.User
	err := us.db.Find(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u, nil
		}
		return nil, err
	}
	return u, nil
}

func (us *UserStore) FindByID(id int) (*model.User, error) {
	var u *model.User
	err := us.db.Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (us *UserStore) FindByUid(uid string) (*model.User, error) {
	var u *model.User
	err := us.db.Where("firebase_uid = ?", uid).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (us *UserStore) Register(user *model.User) error {
	return us.db.Create(user).Error
}

func (us *UserStore) Update(user *model.User) error {
	return us.db.Model(user).Updates(user).Error
}
