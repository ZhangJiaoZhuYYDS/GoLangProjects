// @Author zhangjiaozhu 2023/11/25 9:32:00
package model

import "github.com/jinzhu/gorm"

type Information struct {
	gorm.Model
	UserName string
	UserId   string
	IsRead   bool
	TypeID   int
}
