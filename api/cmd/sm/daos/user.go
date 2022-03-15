package daos

import (
  "fmt"
  "github.com/bdarge/api/cmd/sm/config"
  _ "github.com/bdarge/api/cmd/sm/helper"
  "github.com/bdarge/api/cmd/sm/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
  return &UserDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *UserDAO) Get(id uint) (*models.User, error) {
  var user models.User
  fmt.Printf("get user with id, %s\n", id)

  err := config.Config.DB.Where("id = ?", id).
    First(&user).
    Error

  return &user, err
}

func (dao *UserDAO) GetByAccountId(id uint) (*models.User, error) {
  var user models.User
  fmt.Printf("get user with id, %d\n", id)

  err := config.Config.DB.Where("account_id = ?", id).
    First(&user).
    Error

  return &user, err
}

func (dao *UserDAO) Post(user *models.User) (*models.User, error)  {
  if err := config.Config.DB.Create(&user).Error; err!= nil {
    return nil, err
  }
  return user, nil
}

func (dao *UserDAO) Patch(user *models.User) (*models.User, error)  {
  err := config.Config.DB.Save(&user).
    Error

  return user, err
}
