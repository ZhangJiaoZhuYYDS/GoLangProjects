// @Author zhangjiaozhu 2023/11/25 9:38:00
package model

import "github.com/jinzhu/gorm"

type Type struct {
	gorm.Model
	TypeID   uint
	TypeName string
}
