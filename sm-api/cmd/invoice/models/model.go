package models

import (
	"time"
)

// Login
type Login struct {
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
}

type Register struct {
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
}

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        int       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}

// User Model
type User struct {
  Model
  Email    string  `gorm:"column:email" json:"email"`
  UserName  string `gorm:"column:userName"`
  Password string `gorm:"column:password" json:"-"`
  Orders   []Order `gorm:"foreignKey:CreatedBy"`
  Quotes   []Quote `gorm:"foreignKey:CreatedBy"`
}

// Customer Model
type Customer struct {
	Model
	Email string `gorm:"column:email" json:"email"`
	Name  string `gorm:"column:name" json:"name"`
	Orders []Order `gorm:"null" json:"orders"`
	Quotes []Quote `gorm:"null" json:"quotes"`
}

// Order Model
type Order struct {
	Model
	Description string `gorm:"column:description" json:"description"`
	DeliveryDate  string `gorm:"column:delivery_date" json:"deliveryDate"`
	InvoiceNumber   string `gorm:"column:invoice_number" json:"invoiceNumber"`
	Currency     string `gorm:"column:currency" json:"currency"`
	OrderNumber string `gorm:"column:order_number" json:"orderNumber"`
	Items []OrderItem `json:"items"`
	CreatedBy uint `json:"createdBy"`
	CustomerID uint `json:"customerId"`
}

// Quote Model
type Quote struct {
	Model
	Description string `gorm:"column:description" json:"description"`
	QuoteNumber   string `gorm:"column:quote_number" json:"quoteNumber"`
	Items []QuoteItem `json:"items"`
	CreatedBy uint `json:"createdBy"`
	CustomerID uint `json:"customerId"`
}

// OrderItem Model
type OrderItem struct {
	Model
	Description string `gorm:"column:description" json:"description"`
	Qty  string `gorm:"column:qty" json:"qty"`
	Unit   string `gorm:"column:unit" json:"unit"`
	UnitPrice     string `gorm:"column:unit_price" json:"unitPrice"`
	OrderID uint `json:"orderId"`
}

// QuoteItem Model
type QuoteItem struct {
	Model
	Description string `gorm:"column:description" json:"description"`
	Qty  string `gorm:"column:qty" json:"qty"`
	Unit   string `gorm:"column:unit" json:"unit"`
	UnitPrice     string `gorm:"column:unit_price" json:"unitPrice"`
	QuoteID uint `json:"quoteId"`
}

type Sheet struct {
	Page int `json:"page"`
	Size int `json:"size"`
	TotalElements int64 `json:"totalElements"`
}

type CustomerInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type OrderSheetContent struct {
	Model
	Order
	Customer CustomerInfo `json:"customer"`
}

type QuoteSheetContent struct {
	Model
	Quote
	Customer CustomerInfo `json:"customer"`
}

type CustomerSheet struct {
	Sheet
	Content []Customer `json:"content"`
}

type OrderSheet struct {
	Sheet
	Content []OrderSheetContent `json:"content"`
}

type QuoteSheet struct {
	Sheet
	Content []QuoteSheetContent `json:"content"`
}

type Query struct {
	Page int `form:"page"`
	Size int `form:"size"`
	SortProperty string `form:"sortProperty"`
	SortDirection string `form:"sortDirection"`
	Search string `form:"search"`
}

