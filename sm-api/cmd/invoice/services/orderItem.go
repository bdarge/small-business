package services

import m "github.com/bdarge/sm-api/cmd/invoice/models"

type orderItemDAO interface {
	Get(id uint) (*m.OrderItem, error)
	GetAll(orderId uint) (*[]m.OrderItem, error)
	Post(order *m.OrderItem) (*m.OrderItem, error)
	Delete(id uint) error
	Patch(order *m.OrderItem) (*m.OrderItem, error)
}

type OrderItemService struct {
	dao orderItemDAO
}

func NewOrderItemService(dao orderItemDAO) *OrderItemService {
	return &OrderItemService{dao}
}

func (service *OrderItemService) GetAll(orderId uint) (*[]m.OrderItem, error) {
	return service.dao.GetAll(orderId)
}

func (service *OrderItemService) Post(order *m.OrderItem) (*m.OrderItem, error) {
	return service.dao.Post(order)
}

func (service *OrderItemService) Delete(id uint) error {
	return service.dao.Delete(id)
}

func (service *OrderItemService) Patch(orderItem *m.OrderItem) (*m.OrderItem, error) {
	return service.dao.Patch(orderItem)
}