package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	dsn := "root:Root-123456@tcp(127.0.0.1:3306)/xkshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阀值
			LogLevel:      logger.Info,
			Colorful:      true, //禁用彩色打印
		},
	)
	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不让生成的表名加s
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

}
