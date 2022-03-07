package apis

import (
	"github.com/bdarge/api/cmd/sm/daos"
	"github.com/bdarge/api/cmd/sm/models"
	"github.com/bdarge/api/cmd/sm/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetOrder godoc
// @Summary Retrieves order based on given ID
// @Produce json
// @Param id path integer true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
// @Security ApiKeyAuth
func GetOrder(c *gin.Context) {
	s := services.NewOrderService(daos.NewOrderDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// PostOrder godoc
// @Summary Creates an order
// @Param user body models.Order true "Add order"
// @Success 200 {object} models.Order
// @Router /Orders [post]
// @Security ApiKeyAuth
func PostOrder(c *gin.Context) {
	service := services.NewOrderService(daos.NewOrderDAO())
	order := &models.Order{}

	if err := c.Bind(order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		log.Println("new order => ", order)
		if result, err := service.Post(order); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetOrders godoc
// @Summary Retrieves orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
// @Security ApiKeyAuth
func GetOrders(c *gin.Context) {
	service := services.NewOrderService(daos.NewOrderDAO())
	var query = models.Query{}
	if err := c.BindQuery(&query); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("query %s", query)
		if orderPage, err := service.GetAll(query); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, orderPage)
		}
	}
}

// DeleteOrder godoc
// @Summary Delete order based on given ID
// @Produce json
// @Param id path integer true "Order ID"
// @Success 204
// @Router /orders/{id} [delete]
// @Security ApiKeyAuth
func DeleteOrder(c *gin.Context) {
	service := services.NewOrderService(daos.NewOrderDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := service.Delete(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}

// PatchOrder godoc
// @Summary Delete order based on given ID
// @Produce json
// @Param id path integer true "Order ID"
// @Success 204 {object} models.Order
// @Router /orders/{id} [patch]
// @Security ApiKeyAuth
func PatchOrder(c *gin.Context) {
	s := services.NewOrderService(daos.NewOrderDAO())
	id, _ := strconv.Atoi(c.Param("id"))
	order := &models.Order{}
	if err := c.Bind(order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		order.ID = id
		if result, err := s.Patch(order); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
