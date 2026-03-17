package global

import (
	"fast_gin/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const Version = "0.0.1"

// 全局变量
var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
