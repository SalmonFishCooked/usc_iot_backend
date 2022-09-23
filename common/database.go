package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"usc_iot_backend/model"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	host := "43.138.155.75"
	port := "3306"
	user := "camphers"
	password := "wwteam1234"
	dbname := "iot_db1"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
		charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	//数据库建表
	db.AutoMigrate(&model.Device{})
	db.AutoMigrate(&model.Sensor{})
	db.AutoMigrate(&model.History{})
	db.AutoMigrate(&model.Actuator{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
