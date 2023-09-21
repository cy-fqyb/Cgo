package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	ConfigViper *viper.Viper
	DB          *gorm.DB
	Rdb         *redis.Client
)
