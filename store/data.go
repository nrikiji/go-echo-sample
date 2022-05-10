package store

import (
	"gorm.io/gorm"
)

type DataStore struct {
	db *gorm.DB
}

func NewDataStore(db *gorm.DB) *DataStore {
	return &DataStore{db: db}
}
