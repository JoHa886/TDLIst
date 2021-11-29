package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func DataBase(connString string) {

	//配置日志
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // 慢 SQL 阈值
	// 		LogLevel:                  logger.Silent, // 日志级别
	// 		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  false,         // 禁用彩色打印
	// 	},
	// )

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		fmt.Printf("数据库连接错误：%s.\n", err)
	}
	fmt.Printf("数据库连接成功！")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GlobalDB = db
	migtation()
}
