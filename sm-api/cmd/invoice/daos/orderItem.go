package daos

import (
	."github.com/bdarge/sm-api/cmd/invoice/config"
	m "github.com/bdarge/sm-api/cmd/invoice/models"
)

// OrderItemDAO persists user data in database
type OrderItemDAO struct{}

// NewOrderItemDAO creates a new NewOrderItemDAO
func NewOrderItemDAO() *OrderItemDAO {
	return &OrderItemDAO{}
}

func (dao OrderItemDAO) Get(id uint) (*m.OrderItem, error) {
	var items m.OrderItem

	err := Config.DB.
		Model(&m.OrderItem{}).
		Where("id = ?", id).
		Find(&items).Error

	return &items, err
}

func (dao OrderItemDAO) GetAll(id uint) (*[]m.OrderItem, error) {
	var items []m.OrderItem

	err := Config.DB.
		Model(&m.OrderItem{}).
		Where("order_id = ?", id).
		Find(&items).Error

	return &items, err
}

func (dao OrderItemDAO) Post(orderItem *m.OrderItem) (*m.OrderItem, error) {
	err := Config.DB.Create(&orderItem).Error
	return orderItem, err
}

func (dao OrderItemDAO) Delete(id uint) error {
	err := Config.DB.Delete(&m.OrderItem{}, id).Error
	return err
}

func (dao OrderItemDAO) Patch(orderItem *m.OrderItem) (*m.OrderItem, error) {
	err := Config.DB.UpdateColumns(&orderItem).
		Error

	return orderItem, err
}