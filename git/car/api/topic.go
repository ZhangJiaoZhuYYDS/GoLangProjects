// @Author zhangjiaozhu 2023/12/19 21:19:00
package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// topic 话题

// GetTopic 获取话题列表
func GetTopic(c *gin.Context) {
	service := service.TopicInfoShow{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// CreateCategory 创建话题分类
func CreateCategory(c *gin.Context) {
	service := service.CreateCategoryService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
