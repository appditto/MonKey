package models

import "github.com/jackc/pgtype"

type Stats struct {
	Base
	IPAddress  string         `json:"ip_address" gorm:"not null;index:unique_ban_ip,unique"`
	BanAddress string         `json:"ban_address" gorm:"not null;index:unique_ban_ip,unique"`
	Service    *string        `json:"service" gorm:"index:unique_ban_ip,unique"`
	Count      pgtype.Numeric `json:"count" gorm:"type:numeric;default:0;not null"`
}
