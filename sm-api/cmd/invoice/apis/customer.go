package apis

import (
	"github.com/bdarge/sm-api/cmd/invoice/daos"
	"github.com/bdarge/sm-api/cmd/invoice/models"
	"github.com/bdarge/sm-api/cmd/invoice/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetCustomer godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "Customer ID"
// @Success 200 {object} models.User
// @Router /customers/{id} [get]
// @Security ApiKeyAuth
func GetCustomer(c *gin.Context) {
	s := services.NewCustomerService(daos.NewCustomerDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetCustomers godoc
// @Summary Retrieves customers
// @Produce json
// @Success 200 [] models.Customer
// @Router /customers [get]
// @Security ApiKeyAuth
func GetCustomers(c *gin.Context) {
	s := services.NewCustomerService(daos.NewCustomerDAO())
	var query = models.Query{}
	if err := c.BindQuery(&query); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if (models.Query{}) == query  {
			log.Printf("query %s is empty", query)
			q := c.Request.URL.Query()
			var name = q.Get("name")
			if customers, err := s.GetByName(name); err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				log.Println(err)
			} else {
				c.JSON(http.StatusOK, customers)
			}
		} else {
			log.Printf("query %s", query)
			if customerPage, err := s.GetPaged(query); err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				log.Println(err)
			} else {
				c.JSON(http.StatusOK, customerPage)
			}
		}
	}
}

// PostCustomer godoc
// @Summary Creates a customer
// @Param user body models.Register true "Add customer"
// @Success 200 {object} models.Customer
// @Router /Customers [post]
// @Security ApiKeyAuth
func PostCustomer(c *gin.Context) {
	s := services.NewCustomerService(daos.NewCustomerDAO())
	customer := &models.Customer{}
	if err := c.Bind(customer);  err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		customer := &models.Customer{
			Name:  customer.Name,
			Email: customer.Email,
			Orders: []models.Order{},
			Quotes: []models.Quote{},
		}

		if cust, err := s.Post(customer); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, cust)
		}
	}
}

// DeleteCustomer godoc
// @Summary Delete customer based on given ID
// @Produce json
// @Param id path integer true "Customer ID"
// @Success 204 {object}
// @Router /customers/{id} [get]
// @Security ApiKeyAuth
func DeleteCustomer(c *gin.Context) {
	service := services.NewCustomerService(daos.NewCustomerDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := service.Delete(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}

// PatchCustomer godoc
// @Summary Delete customer based on given ID
// @Produce json
// @Param id path integer true "Customer ID"
// @Success 204 {object}
// @Router /customers/{id} [patch]
// @Security ApiKeyAuth
func PatchCustomer(c *gin.Context) {
	s := services.NewCustomerService(daos.NewCustomerDAO())
	id, _ := strconv.Atoi(c.Param("id"))
	customer := &models.Customer{}
	if err := c.Bind(customer);  err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		customer := &models.Customer{
			Name:  customer.Name,
			Email: customer.Email,
			Model: models.Model{
				ID: id,
			},
		}

		if cust, err := s.Patch(customer); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, cust)
		}
	}
}

