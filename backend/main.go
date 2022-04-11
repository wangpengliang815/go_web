package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

var (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	database = "go_db"
	password = "wpl19950815"
)

func main() {
	// 创建一个默认路由
	r := gin.Default()
	connStr := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	db, err := gorm.Open("mssql", connStr)
	if err != nil {
		log.Fatal("连接数据库失败:" + err.Error())
	}

	db.AutoMigrate(&Person{})
	db.Create(&Person{Name: "w1", Address: "beijing", Age: 18})
	db.Create(&Person{Name: "w2", Address: "shanghai", Age: 18})
	db.Create(&Person{Name: "w3", Address: "shenzhen", Age: 18})

	defer db.Close()
	r.GET("/", func(c *gin.Context) {
		var person []Person
		db.Find(&person)
		c.JSON(http.StatusOK, person)
	})

	r.Run()
}
