package model

import "gorm.io/gorm"

// Actuator 执行器结构体
type Actuator struct {
	gorm.Model
	DeviceID         uint
	Name             string `gorm:"type: varchar(15);not null"`
	ApiTag           string `gorm:"type: varchar(30);not null;unique"`
	Type             int
	TransmissionType int
	DataType         int
	OperationType    int
	SensorType       string
	SerialNumber     string
	Channel          string
	SlaveAddress     string
	FunctionNumber   string
	DataAddress      string
	DataLength       string
	SampleTime       int64
	SampleFormula    string
}
