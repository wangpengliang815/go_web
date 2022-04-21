package api

import (
	"github.com/gin-gonic/gin"
	. "go_web_server/common"
	. "go_web_server/entity"
	"net/http"
	"time"
)

var db = CreateDbConn()

func GetUserToken(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		return
	}
	db.Where("UserName=? And PassWord=?", user.UserName, user.PassWord).First(&user)
	if user.Id != 0 {
		// 执行操作,成功返回Token
		c.JSON(http.StatusOK, ApiResponse.ok)
	}
	c.JSON(http.StatusNotFound, gin.H{"token": time.Now()})
}
