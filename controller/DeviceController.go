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
	json := make(map[string]interface{}) //注意该结构接受的内容
	ctx.ShouldBind(&json)
	if _, ok := json["id"]; !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "传入的值有误"})
		return
	}

	//查询设备
	var device []model.Device
	db.Where("id = ?", json["id"]).Find(&device)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "查询成功",
		"data": device,
	})
}
