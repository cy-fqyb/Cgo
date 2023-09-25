package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Logger:日志组件
// ConfigViper:用来读取配置文件
// DB:使用gorm来连接数据库
// Rdb:连接Redis
var (
	Logger      *zap.SugaredLogger
	ConfigViper *viper.Viper
	DB          *gorm.DB
	Rdb         *redis.Client
)
