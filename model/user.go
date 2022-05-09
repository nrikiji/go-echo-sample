package model

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FirebaseUid string    `json:"-"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
