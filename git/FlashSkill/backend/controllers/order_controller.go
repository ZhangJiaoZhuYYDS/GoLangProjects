// @Author zhangjiaozhu 2023/10/15 13:46:00
package controllers

import (
	"FlashSkill/backend/models"
	"FlashSkill/backend/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

var OrderService  = services.NewOrderService()

func GetAllOrderInfo(ctx *gin.Context)  {
	result, err := OrderService.GetAllorderInfo()
	if err != nil {
		ctx.JSON(200,"服务器繁忙")
		return
	}
	ctx.JSON(200,result)
}

func UpdateOrder(c *gin.Context)  {
	var orderDTO = &models.Order{}
	if err := c.ShouldBind(orderDTO);err != nil {
		c.JSON(200,"格式错误")
		return
	}
	if err := OrderService.Updateorder(orderDTO);err != nil {
		c.JSON(200,"服务器繁忙")
		return
	}
	c.JSON(200,"success")
}

func CreateOrder(c *gin.Context)  {
	//TODO  c.get(user) 从context获取下单用户id
	var orderDto =  &models.Order{}
	err := c.ShouldBindJSON(orderDto)
	if err != nil {
		//TODO 自定义返回状态码
		c.JSON(200,"请求格式错误")
		return
	}
	value, exists := c.Get("userID")
	if !exists {
		c.JSON(200,"用户未登录")
		return
	}
	userID := value.(int64)
	//TODO 1 判断商品库存 --->  库存减一 --> 创建订单 --> 写入数据库
	// 新建rabbitmq 消息
	message := models.NewMessage(orderDto.ProductId,userID)

}
func DelOrder(c *gin.Context)  {
	param := c.Param("id")
	id, _ := strconv.ParseInt(param,10,64)
	if ok := OrderService.DeleteorderByID(id); !ok {
		c.JSON(200,"服务器异常")
		return
	}
	c.JSON(200,"success")
}