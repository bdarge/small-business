package daos

import (
	"fmt"
	. "github.com/bdarge/api/cmd/sm/config"
	"github.com/bdarge/api/cmd/sm/models"
	"github.com/bdarge/api/cmd/sm/util"
	"log"
	"strings"
)

// CustomerDAO persists user data in database
type CustomerDAO struct{}

// NewCustomerDAO creates a new CustomerDAO
func NewCustomerDAO() *CustomerDAO {
	return &CustomerDAO{}
}

// Get does the actual query to database, if customer with specified id is not found error is returned
func (dao *CustomerDAO) Get(id uint) (*models.Customer, error) {
	var customer models.Customer
	fmt.Printf("get customer with id, %s\n", id)

	err := Config.DB.Where("id = ?", id).
		First(&customer).
		Error

	return &customer, err
}

// GetPaged Get paged customers
func (dao *CustomerDAO) GetPaged(query models.Query) (*models.CustomerSheet, error) {
	var customers []models.Customer
	log.Printf("get customers with or without filters; filters=%s; sort porperty=%s\n", query.Search, query.SortProperty)

	if query.Size == 0 {
		query.Size = 10
	}

	if query.SortDirection == "" {
		query.SortDirection = "desc"
	} else {
		query.SortDirection = strings.ToLower(query.SortDirection)
	}

	if query.SortProperty == "" {
		query.SortProperty = "id"
	} else {
		query.SortProperty = util.ToSnakeCase(query.SortProperty)
	}

	err := Config.DB.Limit(query.Size).Offset(query.Page*query.Size).
		Where("true = ?", query.Search == "").
		Or("Name LIKE ?", "%"+query.Search+"%").
		Or("Email LIKE ?", "%"+query.Search+"%").
		Order(query.SortProperty + " " + query.SortDirection).
		Find(&customers).
		Error
	if err != nil {
		return nil, err
	}

	var total int64

	Config.DB.Model(&models.Customer{}).Count(&total)

	page := models.CustomerSheet{
		Content: customers,
		Sheet: models.Sheet{
			Page:          query.Page,
			Size:          query.Size,
			TotalElements: total,
		},
	}

	return &page, nil
}

// GetByName Get flat customers
func (dao *CustomerDAO) GetByName(nameSearch string) (*[]models.Customer, error) {

	var customers []models.Customer
	log.Printf("get customers by name substrings, %s\n", nameSearch)
	err := Config.DB.
		Where("Name LIKE ?", "%"+nameSearch+"%").
		Find(&customers).
		Error

	if err != nil {
		return nil, err
	}

	return &customers, nil
}

func (dao *CustomerDAO) Post(customer *models.Customer) (*models.Customer, error)  {
	err := Config.DB.Create(&customer).
		Error

	return customer, err
}

// Delete customer
func (dao *CustomerDAO) Delete(id uint) error {
	err := Config.DB.Delete(&models.Customer{}, id).Error
	return err
}

func (dao *CustomerDAO) Patch(customer *models.Customer) (*models.Customer, error)  {
	err := Config.DB.UpdateColumns(&customer).
		Error

	return customer, err
}