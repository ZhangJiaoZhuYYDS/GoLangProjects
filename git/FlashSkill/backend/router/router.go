// @Author zhangjiaozhu 2023/10/14 20:16:00
package router

import (
	"FlashSkill/backend/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func New() *gin.Engine {
	r := gin.Default()

	dir, _ := os.Getwd()

	r.LoadHTMLGlob(dir+`\`+`backend\web\views\*`)
	r.Static("/kit", dir+"/backend/web/assets/")

	// TODO 用户登录注册
	api := r.Group("/api/v1/")
	{
		api.POST("/login",controllers.UserLogin)
		api.POST("/register",controllers.UserRegister)
	}
	admin := r.Group("/api/v1/admin")
	//TODO 校验用户是否登录
	admin.Use()
	{
		admin.GET("/product/all", controllers.GetAllProduct)
		admin.POST("/product/update", controllers.UpdateProduct)
		admin.POST("/product/add",controllers.AddProduct)
		admin.DELETE("/product/del/:id",controllers.DelProduct)

		admin.GET("/order/all/info", controllers.GetAllOrderInfo)
		admin.POST("/order/update", controllers.UpdateOrder)
		admin.POST("/order/add",controllers.CreateOrder)
		admin.DELETE("/order/del/:id",controllers.DelOrder)
	}
	return r
}
