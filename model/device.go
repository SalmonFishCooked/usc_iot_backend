package model

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Name    string `gorm:"type: varchar(15);not null"`
	IsOline bool
}
