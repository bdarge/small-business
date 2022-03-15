package services

import "github.com/bdarge/api/cmd/sm/models"

type accountDAO interface {
	Get(id uint) (*models.Account, error)
	Post(user *models.Account) (*models.Account, error)
	GetByEmail(email string) (*models.Account, error)
}

type AccountService struct {
	dao accountDAO
}

// NewAccountService creates a new NewAccountService with the given user DAO.
func NewAccountService(dao accountDAO) *AccountService {
	return &AccountService{dao}
}

// Get just retrieves account using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *AccountService) Get(id uint) (*models.Account, error) {
	return s.dao.Get(id)
}

// Post Create a new user
func (s *AccountService) Post(acct *models.Account) (*models.Account, error) {
	return s.dao.Post(acct)
}

// GetByEmail account by email
func (s *AccountService) GetByEmail(email string) (*models.Account, error) {
	return s.dao.GetByEmail(email)
}