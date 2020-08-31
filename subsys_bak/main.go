package main

import (
	"fmt"
	"net/http"
	"subsys/config"
	"subsys/dbs"
	"github.com/gin-gonic/gin"
	"subsys/routers"
)
func init() {
	config.Setup()
	dbs.DBSetup()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	handler := routerSetting()
	addr := fmt.Sprintf(":%d", config.Config.Common.HttpPort)
	readTimeout := config.Config.Common.ReadTimeout
	writeTimeout := config.Config.Common.WriteTimeout
	maxHeaderBytes := 1 << 20
	
	//设置http服务器
	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}
}

func routerSetting() *gin.Engine {
	ginObj := gin.New() //得到一个gin实例
	ginObj.Use(gin.Logger()) //加载插件
	ginObj.Use(gin.Recovery())

	//静态文件处理
	ginObj.Static("/user", config.Config.Path.UserSavePath)
	ginObj.Static("/point", config.Config.Path.PointSavePath)
	ginObj.Static("/log", config.Config.Path.LogSavePath)

	//用户登录注册模块
	ginObj.POST("/register", routers.Register) 
	ginObj.POST("/login", routers.Login)
	
	return ginObj
}
