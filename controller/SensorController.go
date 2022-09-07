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
	json := make(map[string]interface{}) //注意该结构接受的内容
	ctx.ShouldBind(&json)

	if _, ok := json["deviceID"]; !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "传入的值有误"})
		return
	}

	//查询设备，返回结果
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
