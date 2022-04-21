package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go_web_server/common"
	. "go_web_server/entity"
	. "go_web_server/middleware"
	"net/http"
)

var db = CreateDbConn()

func Login(c *gin.Context) {
	var user User
	_ = c.ShouldBind(&user)
	db.Where("UserName=? And PassWord=?", user.UserName, user.PassWord).First(&user)
	if user.Id != 0 {
		claims := &JWTClaims{
			UserID:   user.Id,
			UserName: user.UserName,
			PassWord: user.PassWord,
		}
		singedToken, _ := GenerateJwtToken(claims)
		c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, "", singedToken))
	} else {
		c.JSON(http.StatusOK, NewApiResponse(http.StatusNotFound, "用户不存在", nil))
	}
}

func GetUserInfo(c *gin.Context) {
	token := c.Query("token")
	claims, err := VerifyAction(token)
	if err != nil {
		fmt.Println(err.Error())
	}
	var user User
	db.Where("UserName=? And PassWord=? And Id=?", claims.UserName, claims.PassWord, claims.UserID).First(&user)
	if user.Id != 0 {
		c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, "", user))
	} else {
		c.JSON(http.StatusNotFound, NewApiResponse(http.StatusNotFound, "没有查询到", nil))
	}
}

func GetList(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, "", users))
}
