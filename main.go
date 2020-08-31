package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"subsys/config"
	"subsys/routers"
)
func init() {
	config.Setup()
	//dbs.DBSetup()
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
	server.ListenAndServe()
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
	ginObj.POST("/register", routers.Register) //注册
	ginObj.POST("/login", routers.Login) //登录
	ginObj.POST("/setpwd", routers.SetPasswd) //修改密码功能

	//升级合约模块
	ginObj.POST("/updata_owner", routers.UpdateOwner)
	ginObj.POST("/upgrade_erclog", routers.UpgradeErclog)
	ginObj.POST("/upgrade_erc200", routers.UpgradeErc200)
	ginObj.POST("/upgrade_user", routers.UpgradeUser)

	//获取余额与总发行量
	ginObj.GET("/balance", routers.GetBalance)
	ginObj.GET("/totalsupply", routers.GetTotalSupply)

	//token操作
	ginObj.POST("/mint", routers.TokenMint)
	//ginObj.POST("/burn", routers.TokenBurn)
	ginObj.POST("/transfer", routers.Transfer)

	//log操作
	ginObj.POST("/pushlog", routers.PushLog)
	ginObj.POST("/querylog", routers.QueryLog)
	
	return ginObj
}
