// @Author zhangjiaozhu 2023/11/25 13:25:00
package api

import (
	"car/serializer"
	"car/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"strconv"
)

func UserLogin(c *gin.Context) {
	// 获取session信息
	session := sessions.Default(c)
	status := 200
	// 从session中获取用户id
	userID := session.Get("userID")
	// 从请求中获取token
	code := c.Request.Header.Get("AuthCode")
	var loginService service.UserLoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(userID, code, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
		logging.Info(err)
	}

}
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Successfully",
	})
}

func MsgUserInfo(c *gin.Context) {
	var service service.MessageInfoService
	if err := c.ShouldBind(&service); err != nil {
		res := service.UserInfo(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
		logging.Info(err)
	}
}

// CheckToken 用户详情
func CheckToken(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}

func UserShow(c *gin.Context) {
	var service service.UserInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfo()
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
		logging.Info(err)
	}
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

func BindPhone(c *gin.Context) {
	var service service.ValidPhoneService
	authorization := c.Request.Header.Get("Authorization")
	if err := c.ShouldBind(&service); err == nil {
		res := service.Valid(authorization)
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
		logging.Info(err)
	}

}

// UserGetCode 发送验证码
func UserGetCode(c *gin.Context) {
	var service service.GetCodeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SendMsg()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func BindEmail(c *gin.Context) {
	var service service.VaildEmailService
	authorization := c.Request.Header.Get("Authorization")
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild(authorization)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
