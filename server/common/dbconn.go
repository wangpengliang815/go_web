package common

import (
	"fmt"
	. "go_web_server/entity"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

// CreateDbConn 返回数据库连接
func CreateDbConn() *gorm.DB {
	var (
		server   = "localhost"
		port     = 1433
		user     = "sa"
		database = "go_web"
		password = "wpl19950815"
	)
	// 数据库连接字符串
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{
		//	QueryFields: true, // QueryFields 模式会根据当前 model 的所有字段名称进行 select
	})
	if err != nil {
		log.Fatal("create connString failed!:" + err.Error())
	}

	// 启用自动迁移生成表
	_ = db.AutoMigrate(&User{})
	return db
}
