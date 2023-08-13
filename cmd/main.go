package main

import (
	"github.com/gin-gonic/gin"
	"go-template/global"
	"go-template/initialize"
	"net/http"
)

func main() {
	//进行初始化
	Init()
	r := gin.New()
	r.GET("ping/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	r.Run(":" + global.Config.Project.Port)
}

func Init() {
	//加载配置
	initialize.InitConfig()
	//初始化zap日志管理
	initialize.InitLogger()
	//初始化mysql
	initialize.InitMysql()
}
