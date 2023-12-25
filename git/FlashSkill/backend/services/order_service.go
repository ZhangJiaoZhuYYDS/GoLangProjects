// @Author zhangjiaozhu 2023/10/15 13:34:00
package services

import (
	"FlashSkill/backend/models"
	"FlashSkill/backend/respositories"
)

type IOrderService interface {
	GetOrderByID(int64)(*models.Order,error)
	DeleteorderByID(int64)bool
	Updateorder(*models.Order) error
	Insertorder(*models.Order)(int64,error)
	GetAllorder()([]*models.Order,error)
	GetAllorderInfo()(map[int]map[string]string,error)

}

type OrderService struct {
	OrderResp respositories.IOrderResp
}

func (o *OrderService) GetOrderByID(i int64) (*models.Order, error) {
	//TODO implement me
	return o.OrderResp.SelectByKey(i)
}

func (o OrderService) DeleteorderByID(i int64) bool {
	//TODO implement me
	return o.OrderResp.Delete(i)
}

func (o OrderService) Updateorder(order *models.Order) error {
	//TODO implement me
	return o.OrderResp.Update(order)
}

func (o OrderService) Insertorder(order *models.Order) (int64 ,error) {
	//TODO implement me
	return o.OrderResp.Insert(order)
}

func (o OrderService) GetAllorder() ([]*models.Order, error) {
	//TODO implement me
	return o.OrderResp.SelectAll()
}

func (o OrderService) GetAllorderInfo() (map[int]map[string]string, error) {
	//TODO implement me
	return o.OrderResp.SelectAllWithInfo()
}

func NewOrderService() IOrderService  {
	return &OrderService{OrderResp: respositories.NewOrderManagerResp("order")}
}