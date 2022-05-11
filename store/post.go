package store

import (
	"errors"
	"go-echo-starter/model"

	"gorm.io/gorm"
)

func (us *UserStore) AllPosts() ([]model.Post, error) {
	var p []model.Post
	err := us.db.Preload("User").Find(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return p, nil
		}
		return nil, err
	}
	return p, nil
}

func (us *UserStore) FindPostByID(id int) (*model.Post, error) {
	var p *model.Post
	err := us.db.Where("id = ?", id).First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func (us *UserStore) UpdatePost(post *model.Post) error {
	return us.db.Model(post).Updates(post).Error
}
