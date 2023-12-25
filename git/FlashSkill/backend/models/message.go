// @Author zhangjiaozhu 2023/10/15 22:17:00
package models

type Message struct {
	ProductID int64
	UserID int64
}

func NewMessage(userID int64,productID int64) *Message {
	return &Message{
		ProductID: userID,
		UserID:    productID,
	}
}
