package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	core.InitGorm()

	// 初始化
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server 运行在: %s", addr)
	router.Run(addr)
}
