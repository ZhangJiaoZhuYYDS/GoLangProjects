package router

import (
	"chat-room/api/v1"
	"chat-room/pkg/common/response"
	"chat-room/pkg/global/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)
	// server.Use(gin.Recovery())
	socket := RunSocekt

	group := server.Group("")
	{
		group.GET("/user", v1.GetUserList)          // 根据用户id获取其所有朋友
		group.GET("/user/:uuid", v1.GetUserDetails) //用户详情
		group.GET("/user/name", v1.GetUserOrGroupByName)
		group.POST("/user/register", v1.Register) // 用户注册
		group.POST("/user/login", v1.Login)       // 用户登录
		group.PUT("/user", v1.ModifyUserInfo)     // 修改用户信息

		group.POST("/friend", v1.AddFriend) //添加朋友

		group.GET("/message", v1.GetMessage) // 查询历史信息

		group.GET("/file/:fileName", v1.GetFile) //获取文件
		group.POST("/file", v1.SaveFile)         // 保存文件

		group.GET("/group/:uuid", v1.GetGroup)
		group.POST("/group/:uuid", v1.SaveGroup)
		group.POST("/group/join/:userUuid/:groupUuid", v1.JoinGroup)
		group.GET("/group/user/:uuid", v1.GetGroupUsers)

		group.GET("/socket.io", socket)

	}
	return server
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("HttpError", zap.Any("HttpError", err))
			}
		}()

		c.Next()
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error("gin catch error: ", log.Any("gin catch error: ", r))
			c.JSON(http.StatusOK, response.FailMsg("系统内部错误"))
		}
	}()
	c.Next()
}
