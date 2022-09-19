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
