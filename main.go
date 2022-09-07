package main

import (
	"github.com/gin-gonic/gin"
	"usc_iot_backend/common"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
