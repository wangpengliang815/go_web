package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
	_ "gorm.io/mssql"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id      int `gorm:"primaryKey"`
	Name    string
	Age     int
	Address string
}

func (Person) TableName() string {
	return "Person"
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		db := createDbConn()
		var person []Person
		db.Find(&person)
		c.JSON(http.StatusOK, person)
	})
	r.Run()
}

// 返回数据库连接
func createDbConn() *gorm.DB {
	var (
		server   = "localhost"
		port     = 1433
		user     = "sa"
		database = "go_db"
		password = "wpl19950815"
	)

	// 数据库连接字符串
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	db, err := gorm.Open("mssql", connString)
	if err != nil {
		log.Fatal("create dbconn failed!:" + err.Error())
	}

	// 启用自动迁移
	db.AutoMigrate(&Person{})
	// 创建2条测试数据
	db.Create(&Person{Name: "wangpengliang", Address: "beijing", Age: 18})
	db.Create(&Person{Name: "lizimeng", Address: "shanghai", Age: 18})

	return db
}
