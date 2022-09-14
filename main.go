package main

import (
	"github.com/gin-gonic/gin"
	"usc_iot_backend/common"
	"usc_iot_backend/tcp"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)

	//启动TCP服务器
	go tcp.CreateTCPServer()
	//启动后端服务
	panic(r.Run())
}
