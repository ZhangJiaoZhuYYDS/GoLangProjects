// @Author zhangjiaozhu 2023/12/19 21:20:00
package service

import (
	"car/model"
	"car/pkg/e"
	"car/serializer"
	logging "github.com/sirupsen/logrus"
)

type TopicInfoShow struct {
}

type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
	EnglishName  string `form:"english_name" json:"english_name"`
}

// List 获取话题分类
func (service *TopicInfoShow) List() serializer.Response {
	var Topics []model.Category
	code := e.Success
	if err := model.DB.Find(&Topics).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTopics(Topics),
		Msg:    e.GetMsg(code),
	}
}

// Create 创建话题分类
func (service *CreateCategoryService) Create() serializer.Response {
	category := model.Category{
		CategoryName: service.CategoryName,
		EnglishName:  service.EnglishName,
	}
	code := e.Success
	err := model.DB.Create(&category).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTopic(category),
		Msg:    e.GetMsg(code),
	}
}
