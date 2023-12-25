// @Author zhangjiaozhu 2023/12/19 21:22:00
package serializer

import "car/model"

// 序列化话题

type Topic struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	EnglishName  string `json:"english_name"`
}

// 序列化话题
func BuildTopic(item model.Category) Topic {
	return Topic{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		EnglishName:  item.EnglishName,
	}
}

// 序列化轮播图列表
func BuildTopics(items []model.Category) (carousels []Topic) {
	for _, item := range items {
		carousel := BuildTopic(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
