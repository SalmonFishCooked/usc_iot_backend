package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"usc_iot_backend/common"
	"usc_iot_backend/model"
)

type HistoryDetail struct {
	ID           uint
	CreatedAt    time.Time
	SensorID     uint
	SensorName   string
	SensorApiTag string
	SensorValue  string
	DeviceApiTag string
}

// GetHistoryInfo 获取设备传感器列表
func GetHistoryInfo(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	type JSON struct {
		DeviceID     int
		StartTime    string
		EndTime      string
		SensorApiTag string
		Page         int
		PageSize     int
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
	var histories []HistoryDetail

	db = db.Table("histories").
		Select("histories.id, histories.created_at, sensors.id as SensorID, sensors.name as SensorName, sensors.api_tag as SensorApiTag, histories.sensor_value, devices.api_tag as DeviceApiTag").
		Joins("JOIN sensors on sensors.id = histories.sensor_id").
		Joins("JOIN devices on devices.id = sensors.device_id")

	//where
	db.Where("devices.id = ?", json.DeviceID)

	if len(json.SensorApiTag) != 0 {
		db = db.Where("sensors.api_tag = ?", json.SensorApiTag)
	}

	err = db.Limit(json.PageSize).Offset(offset).Find(&histories).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "查询成功",
		"data":  histories,
		"total": count,
	})
}

// CreateHistory 创建一条历史传感数据
func CreateHistory(ctx *gin.Context) {
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

// DeleteHistory 删除某条历史记录
func DeleteHistory(ctx *gin.Context) {
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

	//查询设备ID是否存在
	var devices []model.Device
	db.Where("id = ?", json.DeviceID).First(&devices)
	if len(devices) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "不存在这个设备"})
		return
	}

	//查询ApiTag是否存在
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
