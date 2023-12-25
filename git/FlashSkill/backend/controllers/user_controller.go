// @Author zhangjiaozhu 2023/10/15 21:28:00
package controllers

import (
	"FlashSkill/backend/models"
	"FlashSkill/backend/services"
	"github.com/gin-gonic/gin"
)

var UserService = services.NewUserService()

func UserLogin(c *gin.Context)  {
	// TODO 参数校验
	userLoginDto := &models.UserLoginDto{}
	err := c.ShouldBind(userLoginDto)
	if err != nil {
		c.JSON(200,"数据格式有误")
		return
	}
	user, err := UserService.Select(userLoginDto.UserName)
	if err != nil {
		c.JSON(200,"服务器繁忙")
		return
	}
	if user == nil {
		c.JSON(200,"用户不存在")
		return
	}
	// TODO 比较密码
	ok, b := UserService.IsPwdOK(userLoginDto.UserName, userLoginDto.Password)
	if  !b {
		c.JSON(200,"密码错误")
		return
	}
	// TODO 返回token 保存状态
	c.JSON(200,gin.H{"token":ok})

}
func UserRegister(c *gin.Context)  {
	// TODO 参数校验
	userLoginDto := &models.UserLoginDto{}
	err := c.ShouldBind(userLoginDto)
	if err != nil {
		c.JSON(200,"数据格式有误")
		return
	}
	user, err := UserService.Select(userLoginDto.UserName)
	if err != nil {
		c.JSON(200,"服务器繁忙")
		return
	}
	if user != nil {
		c.JSON(200,"用户已经存在")
		return
	}
	password, _ := services.GeneratePassword(userLoginDto.Password)
	UserService.AddUser(&models.User{
		NickName:     "昵称",
		UserName:     userLoginDto.UserName,
		HashPassword: string(password),
	})
	c.JSON(200,"success")
}