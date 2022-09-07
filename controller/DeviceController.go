package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDeviceInfo(ctx *gin.Context) {
	//db := common.GetDB()
	//
	//user := model.Device{
	//	Model:   gorm.Model{},
	//	Name:    "test1",
	//	IsOline: false,
	//}
	//db.Create(&user)

	//获取前端传入的参数
	id := ctx.PostForm("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "传入的id无效"})
		return
	}

	//查询设备

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "查询成功",
	})
}
