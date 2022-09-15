package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usc_iot_backend/common"
	"usc_iot_backend/model"
)

// GetSensorInfo 获取设备传感器列表
func GetSensorInfo(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	type JSON struct {
		DeviceID uint
		ApiTag   string
		Type     int
		Page     int
		PageSize int
	}
	var json JSON
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "传入数据有误"})
		return
	}

	//查询数据，返回结果
	offset := (json.Page - 1) * json.PageSize
	var count int64
	var sensors []model.Sensor

	db = db.Where("device_id = ?", json.DeviceID)
	if len(json.ApiTag) != 0 {
		db = db.Where("api_tag = ? ", json.ApiTag)
	}
	if json.Type != -1 {
		db = db.Where("type = ? ", json.Type)
	}

	err = db.Limit(json.PageSize).Offset(offset).Find(&sensors).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "查询成功",
		"data":  sensors,
		"total": count,
	})
}

// CreateSensor 获取设备传感器列表
func CreateSensor(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	var sensor model.Sensor
	err := ctx.BindJSON(&sensor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "传入数据有误"})
		return
	}

	//查询设备ID是否存在
	var devices []model.Device
	db.Where("id = ?", sensor.DeviceID).First(&devices)
	if len(devices) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "不存在这个设备"})
		return
	}

	//新增数据，返回结果
	if err := db.Create(&sensor).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "创建成功",
			"data": sensor,
		})
	}
}

// DeleteSensor 获取设备传感器列表
func DeleteSensor(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	type JSON struct {
		DeviceID uint
		ApiTag   string
	}
	var json JSON
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "传入数据有误"})
		return
	}

	//判断传感器是否存在
	var sensors []model.Sensor
	db.Where("device_id = ? AND api_tag = ?", json.DeviceID, json.ApiTag).First(&sensors)
	if len(sensors) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "不存在这个传感器"})
		return
	} else {
		db.Delete(&sensors[0])
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功"})
		return
	}
}
