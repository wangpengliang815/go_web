package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go_web_server/api"
	. "go_web_server/middleware"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(Cors())

	// 默认启动页
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	r.POST("/api/user/login", Login)

	r.GET("/api/user/info", GetUserInfo)

	r.GET("/api/user/getList", GetList)

	err := r.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
