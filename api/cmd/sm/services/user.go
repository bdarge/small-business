package services

import "github.com/bdarge/api/cmd/sm/models"

type userDAO interface {
  Get(id uint) (*models.User, error)
  Post(user *models.User) (*models.User, error)
  GetByAccountId(id uint) (*models.User, error)
  Patch(user *models.User) (*models.User, error)
}

type UserService struct {
  dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
  return &UserService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *UserService) Get(id uint) (*models.User, error) {
  return s.dao.Get(id)
}

// GetByAccountId user by account id
func (s *UserService) GetByAccountId(id uint) (*models.User, error) {
  return s.dao.GetByAccountId(id)
}

// Post Create a new user
func (s *UserService) Post(user *models.User) (*models.User, error) {
  return s.dao.Post(user)
}

// Patch Updates an existing user
func (s *UserService) Patch(user *models.User) (*models.User, error) {
  return s.dao.Patch(user)
}

