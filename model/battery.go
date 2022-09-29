package model

import "gorm.io/gorm"

// Battery 电池结构体
type Battery struct {
	gorm.Model
	Electricity string
	Voltage     string
	Temperature string
}
