package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var _db *gorm.DB

func InitDB() {
	// 配置MySQL连接参数
	username := "root"
	password := "spln13spln"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gorm"
	//dsn := "root:spln13spln@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印sql
	})
	if err != nil { // 链接数据库失败，需要使用panic处理
		log.Fatalln("db connected error", err)
	}
	err = _db.AutoMigrate(&UserInfo{}, &UserFollow{}, &Comment{}, &Message{}, &Video{}, &UserLike{})
	if err != nil {
		log.Fatalln("db auto-migrate error", err)
	}
	db, err := _db.DB()
	if err != nil {
		log.Fatalln("db connected error", err)
	}
	//db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
}

func GetDB() *gorm.DB {
	return _db.Session(&gorm.Session{
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 缓存预编译命令
	})
}
