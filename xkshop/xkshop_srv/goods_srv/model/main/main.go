package main

import (
	"crypto/md5" //自带
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
	model2 "xkshop/v1/xkshop_srv/goods_srv/model"
)

//md5加密算法
func getMD5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

//gorm自动生成表
func main() {
	dsn := "root:Root-123456@tcp(127.0.0.1:3306)/xkshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阀值
			LogLevel:      logger.Info,
			Colorful:      true, //禁用彩色打印
		},
	)
	//全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	//生成表
	err = db.AutoMigrate(&model2.Category{},
		&model2.Brands{},
		&model2.GoodsCategoryBrand{},
		&model2.Banner{},
		&model2.Goods{},
	)
	if err != nil {
		panic(err)
	}

}
