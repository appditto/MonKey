package models

import "time"

type StatsAddress struct {
	BanAddress string `gorm:"primaryKey"`
}

type StatsClient struct {
	IPAddress string `gorm:"primaryKey"`
}

type StatsServiceAddress struct {
	Service    string `gorm:"primaryKey"`
	BanAddress string `gorm:"primaryKey"`
}

type StatsServiceDay struct {
	Day          time.Time `gorm:"type:date;primaryKey"`
	Service      string    `gorm:"primaryKey"`
	Served       uint64    `gorm:"not null;default:0"`
	NewAddresses uint64    `gorm:"not null;default:0"`
}

type StatsDay struct {
	Day          time.Time `gorm:"type:date;primaryKey"`
	NewAddresses uint64    `gorm:"not null;default:0"`
	NewClients   uint64    `gorm:"not null;default:0"`
}
