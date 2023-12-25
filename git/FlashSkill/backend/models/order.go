// @Author zhangjiaozhu 2023/10/15 12:42:00
package models

type Order struct {
	ID int64 `sql:"ID"`
	UserId int64 `sql:"userID"`
	ProductId int64 `sql:"productID"`
	OrderStatus int64 `sql:"orderStatus"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)

