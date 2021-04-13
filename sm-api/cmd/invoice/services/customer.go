package services

import "github.com/bdarge/sm-api/cmd/invoice/models"

type customerDAO interface {
  Get(id uint) (*models.Customer, error)
  GetPaged(query models.Query) (*models.CustomerSheet, error)
  GetByName(nameSearch string) (*[]models.Customer, error)
  Post(customer *models.Customer) (*models.Customer, error)
  Delete(id uint) error
  Patch(customer *models.Customer) (*models.Customer, error)
}

type CustomerService struct {
  dao customerDAO
}

// NewCustomerService creates a new CustomerService with the given customer DAO.
func NewCustomerService(dao customerDAO) *CustomerService {
  return &CustomerService{dao}
}

// Get just retrieves customer using Customer DAO, here can be additional logic for processing data retrieved by DAOs
func (s *CustomerService) Get(id uint) (*models.Customer, error) {
  return s.dao.Get(id)
}

func (s *CustomerService) GetPaged(query models.Query) (*models.CustomerSheet, error) {
  return s.dao.GetPaged(query)
}

func (s *CustomerService) GetByName(nameSearch string) (*[]models.Customer, error) {
  return s.dao.GetByName(nameSearch)
}

// Create a new customer
func (s *CustomerService) Post(customer *models.Customer) (*models.Customer, error) {
 return s.dao.Post(customer)
}

func (s *CustomerService) Delete(id uint) error {
  return s.dao.Delete(id)
}

func (s *CustomerService) Patch(customer *models.Customer) (*models.Customer, error) {
  return s.dao.Patch(customer)
}
