// @Author zhangjiaozhu 2023/12/19 21:27:00
package api

import (
	"car/pkg/util"
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 获取所有的帖子
func GetAllSocial(c *gin.Context) {
	service := service.SocialInfoShow{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// 帖子详情
func SocialDetail(c *gin.Context) {
	service := service.ShowSocialService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// 获取我的帖子
func GetMySocial(c *gin.Context) {
	service := service.ShowMySocialService{}
	chain, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := service.Show(chain.UserID)
	c.JSON(200, res)
}

// 新建帖子
func CreateSocial(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	service := service.CreateSocialService{}
	chaim, _ := util.ParseToken(c.GetHeader("Authorization"))
	fileSize := fileHeader.Size
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(file, fileSize, chaim.UserID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// 删除帖子
func DeleteSocial(c *gin.Context) {
	service := service.DeleteSocialService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// 更新帖子
func UpdateSocial(c *gin.Context) {
	service := service.UpdateSocialService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// 搜索帖子
func SearchSocial(c *gin.Context) {
	service := service.SearchSocialsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// 创建帖子图片
//func CreateSocialImg(c *gin.Context) {
//	service := service.CreateImgService{}
//	file , fileHeader ,_ := c.Request.FormFile("file")
//	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
//	fileSize := fileHeader.Size
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Create()
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//		logging.Info(err)
//	}
//}

// 帖子图片的详情地址
func ShowSocialImgs(c *gin.Context) {
	service := service.ShowImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
