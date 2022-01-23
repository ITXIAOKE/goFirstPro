package main

import (
	"crypto/md5" //自带
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
	model2 "xkshop/v1/xkshop_srv/goods_srv/model"

	//"github.com/sethvargo/go-password/password"

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

	//手动造用户数据
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	pwd := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd) //生产用的password

	for i := 0; i < 10; i++ {
		user := model2.User{
			NickName: fmt.Sprintf("xiaoke%d", i),
			Mobile:   fmt.Sprintf("1511234567%d", i),
			Password: pwd,
		}
		db.Save(&user) //这里必须要传地址，否则报错，这里如果使用map的话，则user["ID"]不存在
		//var user1=map[string]interface{}{
		//	"nickna me":"xiaoke",
		//	"age":18,
		//}
		//db.Save(&user1)

	}

	//生成表
	err = db.AutoMigrate(&model2.User{})
	//if err != nil {
	//	panic(err)
	//}

	//暴力破解md5值
	//fmt.Println(getMD5("1234"))
	//加盐就避免暴力破解md5---- 随机字符串+用户密码
	//fmt.Println(getMD5("XXXXXX-1234"))
	//
	//res, err := password.Generate(64, 10, 10, false, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf(res)

	//salt, encodedPwd := password.Encode("123456", nil)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	//check := password.Verify("123456", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	//Using custom options
	//options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 16, HashFunction: sha512.New}
	//salt, encodedPwd := password.Encode("123456", options)
	//pwd := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd) //生产用的password
	//fmt.Println(salt)
	//fmt.Println(len(pwd)) //确保长度不能超过100
	//
	//pwdInfo := strings.Split(pwd, "$")
	//fmt.Println(pwdInfo) //使用debug方法查找切割字符串后，数组中是第一个元素
	//check := password.Verify("123456", pwdInfo[2], pwdInfo[3], options)
	//fmt.Println(check) // true

}
