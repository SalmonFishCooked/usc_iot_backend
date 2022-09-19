package model

import "gorm.io/gorm"

// Device 设备结构体
type Device struct {
	gorm.Model
	Name     string `gorm:"type: varchar(15);not null;unique"`
	ApiTag   string `gorm:"type: varchar(30);not null;unique"`
	IsOnline bool
}
