package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"miniTiktok/conf"
	"time"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, // 开启预编译
		SkipDefaultTransaction: true, // 跳过默认事务

	})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	//配置连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接最大生命周期

	// 连接成功
	DB = db
	fmt.Println("数据库连接成功")
}

// 封装事务操作
func Transaction(fn func(*gorm.DB) error) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
