package daos

import (
	"fmt"
	."github.com/bdarge/sm-api/cmd/invoice/config"
	m "github.com/bdarge/sm-api/cmd/invoice/models"
	"github.com/bdarge/sm-api/cmd/invoice/util"
	"log"
	"strings"
)

// OrderDAO persists user data in database
type OrderDAO struct{}


// NewOrderDAO creates a new NewOrderDAO
func NewOrderDAO() *OrderDAO {
	return &OrderDAO{}
}

// Get order
func (dao *OrderDAO) Get(id uint) (*m.Order, error) {
	var order m.Order
	fmt.Printf("get order with id, %s\n", id)

	err := Config.DB.Where("id = ?", id).
		First(&order).
		Error

	return &order, err
}

// GetAll Get paged orders
func (dao *OrderDAO) GetAll(query m.Query) (*m.OrderSheet, error) {
	var orders = make([]m.OrderSheetContent, 0)
	fmt.Printf("get orders with or without filterl filter = %s; sort porperty = %s\n", query.Search, query.SortProperty)

	if query.Size == 0 {
		query.Size = 10
	}

	if query.SortDirection == "" {
		query.SortDirection = "desc"
	} else {
		query.SortDirection = strings.ToLower(query.SortDirection)
	}

	if query.SortProperty == "" {
		query.SortProperty = "orders.id"
	} else {
		query.SortProperty = "orders." + util.ToSnakeCase(query.SortProperty)
	}

	rows, err := Config.DB.
		Model(&m.Order{}).
		Select("orders.*, customers.name, customers.email").
		Joins("JOIN customers on customers.id = orders.customer_id").
		Limit(query.Size).
		Offset(query.Page*query.Size).
		Where("true = ?", query.Search == "").
		Or("orders.invoice_number LIKE ?", "%"+query.Search+"%").
		Order(query.SortProperty + " " + query.SortDirection).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sheet m.OrderSheetContent
		var order m.Order
		var customer m.CustomerInfo
		err = Config.DB.ScanRows(rows, &order)
		if err != nil {
			log.Fatal(err)
		}
		err = Config.DB.ScanRows(rows, &customer)
		if err != nil {
			log.Fatal(err)
		}
		sheet.Customer = customer
		sheet.Model = order.Model
		sheet.Order = order
		orders = append(orders, sheet)
	}

	var total int64

	Config.DB.Model(&m.Order{}).Count(&total)

	page := m.OrderSheet{
		Content: orders,
		Sheet: m.Sheet{
			Page:          query.Page,
			Size:          query.Size,
			TotalElements: total,
		},
	}

	return &page, nil
}

// Post Create order
func (dao *OrderDAO) Post(order *m.Order) (*m.Order, error)  {
	err := Config.DB.Create(&order).
		Error

	return order, err
}

// Delete order
func (dao *OrderDAO) Delete(id uint) error {
	err := Config.DB.Delete(&m.Order{}, id).Error
	return err
}

// Patch order
func (dao *OrderDAO) Patch(order *m.Order) (*m.Order, error)  {
	err := Config.DB.UpdateColumns(&order).
		Error

	return order, err
}