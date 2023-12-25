// @Author zhangjiaozhu 2023/11/25 13:19:00
package routers

import (
	"car/api"
	"car/routers/middleware"
	"car/service/ws"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// 初始化配置文件
	//conf.Init()
	r := gin.Default()
	// 使用 CORS 中间件
	r.Use(middleware.Cors())
	// 使用session
	store := cookie.NewStore([]byte("cookie_secret"))
	r.Use(sessions.Sessions("mySession", store))

	//
	v1 := r.Group("api/v1")
	{
		v1.GET("/pingTest", api.Ping)
		v1.GET("/MessageIndex/:id", ws.MessageIndex) // 获取最新信息列表
		v1.POST("/user/login", api.UserLogin)        // 用户登录
		v1.GET("/user/get-code", api.UserGetCode)    // 发送腾讯短信验证码
		v1.GET("/topic", api.GetTopic)               // 获取全部话题
		v1.GET("/social", api.GetAllSocial)          // 获取全部帖子
		v1.POST("/upload", api.UpLoad)               // 文件上传到七牛云
		v1.GET("/ws", ws.WsHandler)                  // ws通信
		v1.GET("/get-user-id/:id", api.MsgUserInfo)  // 获取聊天好友的信息
		CommentGroup := v1.Group("/comment/")
		{
			CommentGroup.GET("get-single/:id", api.ShowSingleComm)       //根据id查询单条评论
			CommentGroup.GET("get-children/:id", api.ShowSingleChildren) //根据id查询单条评论的所有子评论
			CommentGroup.GET("get-all/", api.ShowAllComment)             //根据分页条件查找所有的评论
			CommentGroup.GET("children/index", api.ShowAllComChildren)   //
			CommentGroup.POST("create-comment", api.CreateComment)       //新建一条评论
		}
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //用户操作需要jwt验证是否登录状态
		{
			authed.GET("ping", api.CheckToken) //验证token

			//用户操作
			authed.GET("user/show", api.UserShow)    //用户信息展示
			authed.POST("user/email", api.BindEmail) //
			authed.POST("user/phone", api.BindPhone) //用户绑定，解绑手机
			//authed.GET("/get-user-id/:id",api.MessageUserInfo)

			// 好友操作
			authed.GET("friend-all/:id", api.ShowMyFriend) //展示朋友
			authed.GET("friend/:id", api.ShowMyFriendInfo) //展示好友信息
			authed.POST("friend/:id", api.CreateFriend)    //关注好友
			authed.DELETE("friend/:id", api.DeleteFriend)  //删除朋友

			//帖子操作
			authed.POST("social-create", api.CreateSocial)
			//authed.POST("create-social-img",api.CreateSocialImg)
			authed.POST("social-search", api.SearchSocial)
			authed.GET("social-img/:id", api.ShowSocialImgs)
			authed.DELETE("social/:id", api.DeleteSocial)
			authed.GET("social-my", api.GetMySocial)
			authed.GET("social-detail/:id", api.SocialDetail)
		}
	}
	return r
}
