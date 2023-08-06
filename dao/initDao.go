package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"miniTiktok/conf"
)

var DB *gorm.DB

type config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbname string
}

func InitDataBase() {

	conf := &config{
		user:   conf.User,   // 用户名
		pass:   conf.Pass,   // 密码
		adrr:   conf.Adrr,   // 地址
		port:   conf.Port,   // 端口
		dbname: conf.Dbname, // 数据库名称
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
	fmt.Println("数据库连接成功")
}
