package services

import m "github.com/bdarge/api/cmd/sm/models"

type quoteItemDAO interface {
	Get(id uint) (*m.QuoteItem, error)
	GetAll(orderId uint) (*[]m.QuoteItem, error)
	Post(quote *m.QuoteItem) (*m.QuoteItem, error)
	Delete(id uint) error
	Patch(quote *m.QuoteItem) (*m.QuoteItem, error)
}

type QuoteItemService struct {
	dao quoteItemDAO
}

func NewQuoteItemService(dao quoteItemDAO) *QuoteItemService {
	return &QuoteItemService{dao}
}

func (service *QuoteItemService) GetAll(orderId uint) (*[]m.QuoteItem, error) {
	return service.dao.GetAll(orderId)
}

func (service *QuoteItemService) Post(quote *m.QuoteItem) (*m.QuoteItem, error) {
	return service.dao.Post(quote)
}

func (service *QuoteItemService) Delete(id uint) error {
	return service.dao.Delete(id)
}

func (service *QuoteItemService) Patch(quote *m.QuoteItem) (*m.QuoteItem, error) {
	return service.dao.Patch(quote)
}