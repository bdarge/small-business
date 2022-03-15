package daos

import (
	"errors"
	"fmt"
	"github.com/bdarge/api/cmd/sm/config"
	"github.com/bdarge/api/cmd/sm/models"
	"gorm.io/gorm"
)

// AccountDAO persists account data in database
type AccountDAO struct{}

// NewAccountDAO creates a new AccountDAO
func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

// Get does the actual query to database, if account with specified id is not found error is returned
func (dao *AccountDAO) Get(id uint) (*models.Account, error) {
	var account models. Account
	fmt.Printf("get account with id, %d\n", id)

	err := config.Config.DB.Where("id = ?", id).
		First(&account).
		Error

	return &account, err
}

// Post new account. Password is hashed in BeforeCreate hook.
func (dao *AccountDAO) Post(account *models.Account) (*models.Account, error)  {
	user := &models.User{
		Account: *account,
	}
	if err := config.Config.DB.Create(&user).Error; err!= nil {
		return nil, err
	}

	return &user.Account, nil
}

func (dao *AccountDAO) GetByEmail(email string) (*models.Account, error) {
	var account models.Account
	if err := config.Config.DB.Where("email = ?", email).First(&account).Error; err!= nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &account, nil
}
