package services

import m "github.com/bdarge/api/cmd/sm/models"

type orderDAO interface {
	Get(id uint) (*m.Order, error)
	GetAll(query m.Query) (*m.OrderSheet, error)
	Post(order *m.Order) (*m.Order, error)
	Delete(id uint) error
	Patch(order *m.Order) (*m.Order, error)
}

type OrderService struct {
	dao orderDAO
}

func NewOrderService(dao orderDAO) *OrderService {
	return &OrderService{dao}
}

func (service *OrderService) Get(id uint) (*m.Order, error) {
	return service.dao.Get(id)
}

func (service *OrderService) GetAll(query m.Query) (*m.OrderSheet, error) {
	return service.dao.GetAll(query)
}

// Create a new order
func (service *OrderService) Post(order *m.Order) (*m.Order, error) {
	return service.dao.Post(order)
}

// Delete an order
func (service *OrderService) Delete(id uint) error {
	return service.dao.Delete(id)
}

// patch
func (service *OrderService) Patch(order *m.Order) (*m.Order, error) {
	return service.dao.Patch(order)
}
