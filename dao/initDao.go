package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbname string
}

func Init() {

	conf := &config{
		user:   "guest",         // 用户名
		pass:   "123456",        // 密码
		adrr:   "39.101.72.240", // 地址
		port:   "3306",          // 端口
		dbname: "tiktok",        // 数据库名称
	}
	// 加载数据库连接的链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.user, conf.pass, conf.adrr, conf.port, conf.dbname)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	// 连接成功
	DB = db
	fmt.Println("success")
}
