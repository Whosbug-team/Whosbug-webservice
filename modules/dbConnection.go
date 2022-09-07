package modules

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
	. "webService_Refactoring/utils"
)

// Db 数据库
var Db *gorm.DB
var err error

// InitDB 初始化数据库
func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DbHost,
		DbUser,
		DbPassword,
		DbName,
		DbPort,
		DbSSLMode,
		DbTimeZone,
	)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为Info, 即打印所有SQL语句
		// Logger:logger.Default.LogMode(logger.Warn), // 只打印慢查询, 默认的SlowThreshold为200ms
	})
	if err != nil {
		// 关闭数据库连接，打印错误信息
		log.Fatalf("数据库连接错误: %v", err)
	}
	Create() // 调用dbMigration中的create, 创建表格
	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalf("数据库初始化错误: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)                  // 设置连接池最大空闲连接数
	sqlDB.SetMaxOpenConns(200)                 // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置最大连接时长, 注意不要大于GIN框架的timeout时间
}
