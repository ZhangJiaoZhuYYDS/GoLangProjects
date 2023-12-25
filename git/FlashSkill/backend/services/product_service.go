// @Author zhangjiaozhu 2023/10/14 19:47:00
package services

import (
	"FlashSkill/backend/models"
	"FlashSkill/backend/respositories"
)

type IProductService interface {
	GetProductByID(int64) (*models.Product,error)
	GetAllProduct()([]*models.Product,error)
	DeleteProductByID(int64)bool
	InsertProduct(product *models.Product)(int64,error)
	UpdateProduct(product *models.Product)error
}
type ProductService struct {
	ProductResp respositories.IProduct
}



func (p *ProductService) GetProductByID(i int64) (*models.Product, error) {
	//TODO implement me
	return p.ProductResp.SelectByKey(i)
}

func (p *ProductService) GetAllProduct() ([]*models.Product, error) {
	//TODO implement me
	return p.ProductResp.SelectAll()
}

func (p *ProductService) DeleteProductByID(i int64) bool {
	//TODO implement me
	return p.ProductResp.Delete(i)
}

func (p *ProductService) InsertProduct(product *models.Product) (int64, error) {
	//TODO implement me
	return p.ProductResp.Insert(product)
}

func (p *ProductService) UpdateProduct(product *models.Product) error {
	//TODO implement me
	return p.ProductResp.Update(product)
}

func NewProductService()IProductService  {
	return &ProductService{ProductResp: respositories.NewProductManager("product")}
}