// @Author zhangjiaozhu 2023/11/25 9:37:00
package model

import "github.com/jinzhu/gorm"

type Report struct {
	gorm.Model
	TypeID   uint
	TypeName string
	UserID   uint
	UserName string
	Content  string
	Picture  string
	Finish   uint // 0表示未完成 1表示已完成
}
