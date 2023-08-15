package main

import (
	"go-template/global"
	"go-template/initialize"
	"go-template/router"
)

func main() {
	//进行初始化
	Init()
	r := router.InitApiRouter()
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
