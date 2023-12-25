// @Author zhangjiaozhu 2023/12/19 21:33:00
package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var service service.UpLoadFile
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpLoadFile(file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
