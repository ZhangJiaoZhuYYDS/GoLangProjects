// @Author zhangjiaozhu 2023/11/25 9:35:00
package model

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	CarName     string //车名字
	CarNum      string //车牌号
	CarImages   string //车照片
	CarBossId   uint   //车主ID
	CarBossName string
}
