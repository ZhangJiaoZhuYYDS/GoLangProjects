package routes

import (
	"CarDemo1/api"
	"CarDemo1/conf"
	"CarDemo1/middleware"
	"CarDemo1/service/ws"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	conf.Init()
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	//middleware.HttpLogToFile(conf.AppMode)
	//r.Use(middleware.LoggerToFile())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.GET("/pingTest", api.Ping)
		v1.GET("/MessageIndex/:id", ws.MessageIndex)    //获取最新信息列表
		v1.GET("/user/get-code", api.UserGetCode)       //获得手机验证码 1
		v1.POST("/user/login", api.UserLogin)           //用户登陆 2
		v1.GET("/topic", api.GetTopic)                  //从话题表中获取全部话题 3
		v1.GET("/social", api.GetAllSocial)             //分页查询获取全部帖子信息 4
		v1.POST("/upload", api.UpLoad)                  //文件上传到七牛云 5
		v1.GET("/ws", ws.WsHandler)                     //ws通信
		v1.GET("/get-user-id/:id", api.MessageUserInfo) //获取聊天好友的信息
		CommentGroup := v1.Group("/comment/")
		{
			CommentGroup.GET("get-single/:id", api.ShowSingleComm)     // 根据id从评论表获取单条评论及其子评论
			CommentGroup.GET("get-children", api.ShowSingleChildren)   // 根据id从评论表获取单条评论及其子评论详情
			CommentGroup.GET("get-all/:id", api.ShowAllComment)        // 分页获取所有评论
			CommentGroup.GET("children/index", api.ShowAllComChildren) //关键词模糊查询一个评论下的所有评论
			CommentGroup.POST("create-comment", api.CreateComment)     //新建评论
		}

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("ping", api.CheckToken) //验证有无token

			// 好友操作
			authed.GET("friend-all/:id", api.ShowMyFriend) //根据用户id获取其朋友列表
			authed.GET("friend/:id", api.ShowMyFriendInfo) //根据朋友id获取其朋友信息
			authed.POST("friend/:id", api.CreateFriend)    //根据用户id和朋友添加朋友
			authed.DELETE("friend/:id", api.DeleteFriend)  //删除好友

			//用户操作
			authed.GET("user/show", api.UserShow)    //展示用户详情
			authed.POST("user/email", api.BindEmail) //绑定邮箱
			authed.POST("user/phone", api.BindPhone) //绑定手机号
			//authed.GET("/get-user-id/:id",api.MessageUserInfo)

			//帖子操作
			authed.POST("social-create", api.CreateSocial) //新建帖子
			//authed.POST("create-social-img",api.CreateSocialImg)
			authed.POST("social-search", api.SearchSocial) //搜索帖子
			authed.GET("social-img/:id", api.ShowSocialImgs)
			authed.DELETE("social/:id", api.DeleteSocial)
			authed.GET("social-my", api.GetMySocial)          //展示我的帖子
			authed.GET("social-detail/:id", api.SocialDetail) //帖子详情

			//车辆操作
			authed.GET("cars", api.ShowCar)
			authed.POST("cars", api.CreateCar)
			authed.DELETE("car/:id", api.DeleteCar)
			authed.POST("car", api.SearchCar)

			//反馈操作
			authed.POST("report", api.CreateReport)
			authed.DELETE("report/:id", api.DeleteReport)
			authed.GET("report", api.ShowReport)

			//authed.POST("upload",api.UpLoad)							//上传操作
			//authed.POST("user/sending-email", api.SendEmail) 			//邮箱发送
			authed.POST("comment/delete-comment", api.DeleteComment) //删除评论
			//authed.GET("/ws", ws.WsHandler)  //通信

			//authed.GET("MessageIndex/:id",ws.MessageIndex) 			//获取最新信息列表
		}
	}

	v2 := r.Group("api/v2")
	{
		//管理员登陆注册
		v2.POST("admin/register", api.AdminRegister)
		//登陆
		v2.POST("admin/login", api.AdminLogin)
		v2.GET("topic", api.GetTopic)             //获得分类
		v2.GET("socials", api.GetAllSocial)       //获得帖子
		v2.GET("socials/:id", api.SocialDetail)   //获取帖子
		v2.GET("carousels", api.ListCarousels)    //获得轮播图
		v2.GET("users", api.ListUsers)            //获取用户列表
		v2.GET("reports", api.ShowReport)         //获取用户列表
		v2.GET("cars", api.ListCars)              //获取车辆
		v2.DELETE("cars/:car_num", api.DeleteCar) //获取车辆
		authed2 := v2.Group("/")
		authed2.Use(middleware.JWTAdmin())
		{
			authed2.POST("carousels", api.CreateCarousel)    //创建轮播图
			authed2.DELETE("products/:id", api.DeleteSocial) //删除帖子
			authed2.POST("/update-social", api.UpdateSocial) //更新帖子
			authed2.POST("categories", api.CreateCategory)   //创建分类
			authed2.PUT("report", api.UpdateReport)          //更新Report
		}
	}
	return r
}
