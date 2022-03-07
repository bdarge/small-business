package daos

import (
	"fmt"
	."github.com/bdarge/api/cmd/sm/config"
	m "github.com/bdarge/api/cmd/sm/models"
	"github.com/bdarge/api/cmd/sm/util"
	"log"
	"strings"
)

// QuoteDAO persists user data in database
type QuoteDAO struct{}


// NewQuoteDAO creates a new NewQuoteDAO
func NewQuoteDAO() *QuoteDAO {
	return &QuoteDAO{}
}

// Get quote
func (dao *QuoteDAO) Get(id uint) (*m.Quote, error) {
	var quote m.Quote
	fmt.Printf("get quote with id, %s\n", id)

	err := Config.DB.Where("id = ?", id).
		First(&quote).
		Error

	return &quote, err
}

// Get paged quotes
func (dao *QuoteDAO) GetAll(query m.Query) (*m.QuoteSheet, error) {
	var quotes = make([]m.QuoteSheetContent,0)

	fmt.Printf("get quotes with or without filters; filter=%s; sort porperty=%s\n", query.Search, query.SortProperty)

	if query.Size == 0 {
		query.Size = 10
	}

	if query.SortDirection == "" {
		query.SortDirection = "desc"
	} else {
		query.SortDirection = strings.ToLower(query.SortDirection)
	}

	if query.SortProperty == "" {
		query.SortProperty = "quotes.id"
	} else {
		query.SortProperty = "quotes." + util.ToSnakeCase(query.SortProperty)
	}

	rows, err := Config.DB.
		Model(&m.Quote{}).
		Select("quotes.*, customers.*").
		Joins("JOIN customers on customers.id = quotes.customer_id").
		Limit(query.Size).
		Offset(query.Page*query.Size).
		Where("true = ?", query.Search == "").
		Or("Name LIKE ?", "%"+query.Search+"%").
		Or("Email LIKE ?", "%"+query.Search+"%").
		Order(query.SortProperty + " " + query.SortDirection).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var quote m.Quote
		var sheet m.QuoteSheetContent
		var customer m.CustomerInfo
		err = Config.DB.ScanRows(rows, &quote)
		if err != nil {
			log.Fatal(err)
		}
		err = Config.DB.ScanRows(rows, &customer)
		if err != nil {
			log.Fatal(err)
		}
		sheet.Customer = customer
		sheet.Quote = quote
		quotes = append(quotes, sheet)
	}

	var total int64

	Config.DB.Model(&m.Quote{}).Count(&total)

	page := m.QuoteSheet{
		Content: quotes,
		Sheet: m.Sheet{
			Page:          query.Page,
			Size:          query.Size,
			TotalElements: total,
		},
	}

	return &page, nil
}

// Create quote
func (dao *QuoteDAO) Post(quote *m.Quote) (*m.Quote, error)  {
	err := Config.DB.Create(&quote).
		Error

	return quote, err
}

// Delete quote
func (dao *QuoteDAO) Delete(id uint) error {
	err := Config.DB.Delete(&m.Quote{}, id).Error
	return err
}

// patch
func (dao *QuoteDAO) Patch(q *m.Quote) (*m.Quote, error)  {
	err := Config.DB.UpdateColumns(&q).
		Error

	return q, err
}