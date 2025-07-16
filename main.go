package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Error("xxx")
	global.Log.Warning("warn")
	global.Log.Info("info")
	// 连接数据库
	core.InitGorm()
}
