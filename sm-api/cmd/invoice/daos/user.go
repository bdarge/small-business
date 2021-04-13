package daos

import (
  "fmt"
  "github.com/bdarge/sm-api/cmd/invoice/config"
  "github.com/bdarge/sm-api/cmd/invoice/httputil"
  "github.com/bdarge/sm-api/cmd/invoice/models"
  _ "github.com/bdarge/sm-api/cmd/invoice/httputil"
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

func (dao *UserDAO) Post(user *models.User) (*models.User, error)  {
  password, e := httputil.HashPassword(user.Password);
  if e != nil {
    return nil, e
  }
  user.Password = password
  err := config.Config.DB.Create(&user).
    Error

  return user, err
}
