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

// PostOrderItem godoc
// @Summary Creates a order
// @Param user body models.Register true "Add orderItem"
// @Success 200 {object} models.OrderItem
// @Router /OrderItems [post]
// @Security ApiKeyAuth
func PostOrderItem(c *gin.Context) {
	service := services.NewOrderItemService(daos.NewOrderItemDAO())
	orderItem := &models.OrderItem{}

	if err := c.Bind(orderItem); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		if result, err := service.Post(orderItem); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetOrderItems godoc
// @Summary Retrieves orders
// @Produce json
// @Success 200 [] models.OrderItem
// @Router /{id}/orderItems [get]
// @Security ApiKeyAuth
func GetOrderItems(c *gin.Context) {
	service := services.NewOrderItemService(daos.NewOrderItemDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if orderItems, err := service.GetAll(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, orderItems)
	}
}

// DeleteOrderItem godoc
// @Summary Delete orderItem based on given ID, and order ID
// @Produce json
// @Param id path integer true "Order ID"
// @Success 204 {object}
// @Router /orderItems/{itemId} [get]
// @Security ApiKeyAuth
func DeleteOrderItem(c *gin.Context) {
	service := services.NewOrderItemService(daos.NewOrderItemDAO())
	itemId, _ := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err := service.Delete(uint(itemId)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}

// PatchOrderItem godoc
// @Summary Delete orderItem based on given ID
// @Produce json
// @Param id path integer true "OrderItem ID"
// @Success 204 {object}
// @Router /orderItem/{id} [patch]
// @Security ApiKeyAuth
func PatchOrderItem(c *gin.Context) {
	s := services.NewOrderItemService(daos.NewOrderItemDAO())
	itemId, _ := strconv.Atoi(c.Param("itemId"))
	orderItem := &models.OrderItem{}
	if err := c.Bind(orderItem);  err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		orderItem.ID = itemId

		if result, err := s.Patch(orderItem); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
