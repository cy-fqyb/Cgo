package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reggie/global"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库连接
func InitDb() *gorm.DB {
	dbConfig := global.ConfigViper.GetStringMapString("mysql")
	if dbConfig == nil {
		return nil
	}
	//构建数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["dbname"])
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	// 连接数据库
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGormLogger(),
	}); err != nil {
		global.Logger.Error("Mysql connect failed", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.ConfigViper.GetInt("mysql.MaxIdleConns"))
		sqlDB.SetMaxOpenConns(global.ConfigViper.GetInt("mysql.MaxOpenConns"))
		return db
	}
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.ConfigViper.GetBool("mysql.EnableFileLogWriter") {
		stSepartor := string(filepath.Separator)
		stRootDir, _ := os.Getwd()
		stLogFilePath := stRootDir + stSepartor + "log" + stSepartor + time.Now().Format(time.DateOnly) + "-sql.log"

		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   stLogFilePath,
			MaxSize:    global.ConfigViper.GetInt("log.MaxSize"),
			MaxAge:     global.ConfigViper.GetInt("log.MaxAge"),
			Compress:   global.ConfigViper.GetBool("log.Compress"),
			MaxBackups: global.ConfigViper.GetInt("log.MaxBackups"),
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// 自定义sql日志
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.ConfigViper.GetString("mysql.logMode") {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
		LogLevel:                  logMode,                                         // 日志级别
		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  global.ConfigViper.GetBool("mysql.logColorful"), // 禁用彩色打印
	})
}
