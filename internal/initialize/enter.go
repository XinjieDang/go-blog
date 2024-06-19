package initialize

import (
	"dxj/configs"
	"dxj/internal/global"
	"dxj/internal/pkg/logger"
	mysqlDB "dxj/internal/pkg/mysql"
	redisDB "dxj/internal/pkg/redis"
)

func GlobalInit() {

	global.Config = configs.EnvConfig
	//日志初始化
	logger.Init()
	//数据库初始化
	global.DB = mysqlDB.Connect(&configs.EnvConfig.Mysql)
	//redis初始化
	global.Redis = redisDB.Connect(&configs.EnvConfig.Redis)
}
