package services

import m "github.com/bdarge/sm-api/cmd/invoice/models"

type quoteDAO interface {
	Get(id uint) (*m.Quote, error)
	GetAll(query m.Query) (*m.QuoteSheet, error)
	Post(quote *m.Quote) (*m.Quote, error)
	Delete(id uint) error
	Patch(quote *m.Quote) (*m.Quote, error)
}

type QuoteService struct {
	dao quoteDAO
}

func NewQuoteService(dao quoteDAO) *QuoteService {
	return &QuoteService{dao}
}

func (service *QuoteService) Get(id uint) (*m.Quote, error) {
	return service.dao.Get(id)
}

func (service *QuoteService) GetAll(query m.Query) (*m.QuoteSheet, error) {
	return service.dao.GetAll(query)
}

// Create a new quote
func (service *QuoteService) Post(quote *m.Quote) (*m.Quote, error) {
	return service.dao.Post(quote)
}

// Delete an quote
func (service *QuoteService) Delete(id uint) error {
	return service.dao.Delete(id)
}

// patch
func (service *QuoteService) Patch(quote *m.Quote) (*m.Quote, error) {
	return service.dao.Patch(quote)
}

