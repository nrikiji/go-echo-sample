package store

import (
	"errors"
	"go-echo-starter/model"

	"gorm.io/gorm"
)

func (ds *DataStore) AllUsers() ([]model.User, error) {
	var u []model.User
	err := ds.db.Find(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u, nil
		}
		return nil, err
	}
	return u, nil
}

func (ds *DataStore) FindUserByID(id int) (*model.User, error) {
	var u *model.User
	err := ds.db.Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (ds *DataStore) FindUserByUid(uid string) (*model.User, error) {
	var u *model.User
	err := ds.db.Where("firebase_uid = ?", uid).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (ds *DataStore) RegisterUser(user *model.User) error {
	return ds.db.Create(user).Error
}

func (ds *DataStore) UpdateUser(user *model.User) error {
	return ds.db.Model(user).Updates(user).Error
}
