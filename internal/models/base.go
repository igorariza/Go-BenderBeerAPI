package models

import (
	"time"

	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *Base) BeforeCreate(tx *gorm.DB) error {
	//u.ID = uuid.NewV4().String()
	return nil

}
