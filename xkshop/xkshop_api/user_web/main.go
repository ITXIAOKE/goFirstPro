package main

import (
	"fmt"
	"go.uber.org/zap"
	"xkshop/v1/xkshop_api/user_web/initialize"
)

func main() {
	port := 8021
	//logger, _ := zap.NewD evelopment()
	//defer logger.Sync()
	//sugar := logger.Sugar()

	//zap.L() //全局的logger
	//zap.S() //全局的global
	//zap.ReplaceGlobals(logger) //让日志打印输出

	//1,初始化log
	initialize.InitLogger()

	//2,.初始化routers
	engine := initialize.Routers()
	//日志相关debug  info warn error fetal
	zap.S().Debugf("启动服务器，端口：%d", port)

	if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}

}
