package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usc_iot_backend/common"
	"usc_iot_backend/model"
)

// GetDeviceInfo 获取某个设备信息
func GetDeviceInfo(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	type JSON struct {
		ID       int
		Page     int
		PageSize int
	}
	var json JSON
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "传入数据有误"})
		return
	}

	//查询设备
	offset := (json.Page - 1) * json.PageSize
	var count int64
	var devices []model.Device

	if json.ID != -1 {
		db = db.Where("id = ?", json.ID)
	}

	err = db.Limit(json.PageSize).Offset(offset).Find(&devices).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		panic(err)
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "查询成功",
		"data":  devices,
		"total": count,
	})
}

// CreateDevice 创建设备
func CreateDevice(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	var device model.Device
	err := ctx.BindJSON(&device)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "传入数据有误"})
		return
	}

	//新增数据，返回结果
	if err := db.Create(&device).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "创建成功",
			"data": device,
		})
	}
}

// DeleteDevice 删除设备
func DeleteDevice(ctx *gin.Context) {
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
