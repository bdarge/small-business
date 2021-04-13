package daos

import (
	."github.com/bdarge/sm-api/cmd/invoice/config"
	m "github.com/bdarge/sm-api/cmd/invoice/models"
)

// QuoteItemDAO persists user data in database
type QuoteItemDAO struct{}

// NewQuoteItemDAO creates a new NewQuoteItemDAO
func NewQuoteItemDAO() *QuoteItemDAO {
	return &QuoteItemDAO{}
}

func (dao QuoteItemDAO) Get(id uint) (*m.QuoteItem, error) {
	var items m.QuoteItem

	err := Config.DB.
		Model(&m.QuoteItem{}).
		Where("id = ?", id).
		Find(&items).Error

	return &items, err
}

func (dao QuoteItemDAO) GetAll(id uint) (*[]m.QuoteItem, error) {
	var items []m.QuoteItem

	err := Config.DB.
		Model(&m.QuoteItem{}).
		Where("quote_id = ?", id).
		Find(&items).Error

	return &items, err
}

func (dao QuoteItemDAO) Post(orderItem *m.QuoteItem) (*m.QuoteItem, error) {
	err := Config.DB.Create(&orderItem).Error
	return orderItem, err
}

func (dao QuoteItemDAO) Delete(id uint) error {
	err := Config.DB.Delete(&m.QuoteItem{}, id).Error
	return err
}

func (dao QuoteItemDAO) Patch(orderItem *m.QuoteItem) (*m.QuoteItem, error) {
	err := Config.DB.UpdateColumns(&orderItem).
		Error

	return orderItem, err
}
