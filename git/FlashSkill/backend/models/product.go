// @Author zhangjiaozhu 2023/10/14 17:18:00
package models

type Product struct {
	ID int64 `json:"id" sql:"ID" skill:"id"`
	ProductNum int64 `json:"productNum" sql:"productNum"`
	ProductName string `json:"productName" sql:"productName"`
	ProductImage string `json:"productImage" sql:"productImage"`
	ProductUrl string  `json:"productUrl" sql:"productUrl"`
}

type ProductUpdateDto struct {
	ProductNum int64 `json:"productNum" sql:"productNum"`
	ProductName string `json:"productName" sql:"productName"`
	ProductImage string `json:"productImage" sql:"productImage"`
	ProductUrl string  `json:"productUrl" sql:"productUrl"`
}