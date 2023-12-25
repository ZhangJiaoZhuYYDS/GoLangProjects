// @Author zhangjiaozhu 2023/11/25 9:36:00
package model

import "github.com/jinzhu/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint
}
