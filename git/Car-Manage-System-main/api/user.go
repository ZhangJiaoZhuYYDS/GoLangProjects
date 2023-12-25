package api

import (
	"CarDemo1/serializer"
	"CarDemo1/service"
	logging "github.com/sirupsen/logrus"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserLogin 用户登陆接口
// 前端点击扫码授权登录按钮获取用户信息，没有从存储器找到用户信息，就携带空token到后端服务器的校验token接口，后端服务器的校验接口中间件发现没有token,就返回前端没有token的代码
// 1 无token就发起请求并携带生成的有效期验证码和用户信息发送给后端服务器 3 后端服务器携带配置文件中的小程序密钥和有效期验证码发送给微信授权服务器 4 检验成功后返回给后端服务器唯一的openid
// 5 服务器拿到openid就把用户信息存入数据库并生成token返回给前端
// 6 前端把token解析出来用户信息并存储在存储器
// 7 下次前端校验到token直接加载用户信息，不用再登录直到token过期失效
func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	userID := session.Get("userId")
	code := c.Request.Header.Get("AuthCode")
	var loginService service.UserLoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(userID, code, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Successfully",
	})
}

func MessageUserInfo(c *gin.Context) {
	var service service.MessageInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfo(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
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

func UserShow(c *gin.Context) {
	var service service.UserInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfo()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
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
	var service service.VaildPhoneService
	//var operationType = c.Request.Header.Get("operation_type")
	authorization := c.Request.Header.Get("Authorization")
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild(authorization)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UserGetCode 生成随机验证码，一份存入redis,一份通过腾讯信息发送到手机用户
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
