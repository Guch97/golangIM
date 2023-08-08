package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitMySQl() {
	//自定义日志模块   打印语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, _ := gorm.Open(mysql.Open("root:Guch123456.@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb3&parseTime=True&loc=Local"), &gorm.Config{Logger: newLogger})
	fmt.Println(db, "dbdbdbdbdbdbdbdbdb")
	DB = db
}
