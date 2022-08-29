package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables
type Base struct {
	ID        uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// uuid.New() creates a new random UUID or panics.
	base.ID = uuid.New()

	// generate timestamps
	now := time.Now().UTC()
	base.CreatedAt = now

	return nil
}
