package model

import "gorm.io/gorm"

// History 历史传感数据结构体
type History struct {
	gorm.Model
	DeviceID    uint
	SensorID    uint
	SensorValue string
}
