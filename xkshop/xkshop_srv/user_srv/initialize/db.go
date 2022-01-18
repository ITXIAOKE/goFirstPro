package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"xkshop/v1/xkshop_srv/user_srv/global"
)

func InitDB() {
	c := global.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)

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
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "xk_",//所有表名的前缀
			SingularTable: true, //不让生成的表名加s，都是单数
		},
		Logger: newLogger,
		//PrepareStmt: true,//预加载
		//CreateBatchSize: 1000,//这个是批量创建，每次1000，分几批创建完成，比如10000，分10批，每批1000
		//SkipDefaultTransaction: true,//关闭默认事务

	})
	if err != nil {
		panic(err)
	}
}
