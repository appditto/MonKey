package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stats struct {
	ID         uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	IPAddress  string    `json:"ip_address" gorm:"not null;index:unique_ban_ip,unique"`
	BanAddress string    `json:"ban_address" gorm:"not null;index:unique_ban_ip,unique"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;index:unique_ban_ip,unique"`
	Service    *string   `json:"service" gorm:"index:unique_ban_ip,unique"`
	Count      uint64    `json:"count" gorm:"type:numeric;default:0;not null"`
}

// BeforeCreate will set Base struct before every insert
func (stats *Stats) BeforeCreate(tx *gorm.DB) error {
	// uuid.New() creates a new random UUID or panics.
	stats.ID = uuid.New()

	// generate timestamps
	now := time.Now().UTC()
	stats.CreatedAt = now

	return nil
}
