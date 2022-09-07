package model

import "gorm.io/gorm"

// Sensor 传感器结构体
type Sensor struct {
	gorm.Model
	DeviceID         uint
	Name             string `gorm:"type: varchar(15);not null"`
	ApiTag           string `gorm:"type: varchar(30);not null;unique"`
	Type             int
	TransmissionType int
	DataType         int
	Unit             string `gorm:"type: varchar(15);"`
}
