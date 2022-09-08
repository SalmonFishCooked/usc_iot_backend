package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"usc_iot_backend/common"
	"usc_iot_backend/model"
)

// GetSensorInfo 获取设备传感器列表
func GetSensorInfo(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	json := make(map[string]interface{}) //注意该结构接受的内容
	ctx.ShouldBind(&json)

	if _, ok := json["deviceID"]; !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "传入的值有误"})
		return
	}

	//查询数据，返回结果
	var sensors []model.Sensor
	if _, ok := json["apiTag"]; !ok {
		//1.不传入apiTag
		db.Where("device_id = ?", json["deviceID"]).Find(&sensors)
	} else {
		//2.传入apiTag
		db.Where("device_id = ? AND api_tag = ? ", json["deviceID"], json["apiTag"]).First(&sensors)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": sensors,
	})
}

func CreateSensor(ctx *gin.Context) {
	db := common.GetDB()

	//获取前端传入的参数
	var sensor model.Sensor
	err := ctx.BindJSON(&sensor)
	fmt.Println(sensor)
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
			"msg":  "查询成功",
			"data": sensor,
		})
	}
}
