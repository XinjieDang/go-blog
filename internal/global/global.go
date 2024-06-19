package global

import (
	"dxj/configs"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *configs.Config // 全局Config
	DB     *gorm.DB
	Redis  *redis.Client
	Api    *configs.Api
)
