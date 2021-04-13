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

// PostQuoteItem godoc
// @Summary Creates a order
// @Param user body models.Register true "Add quoteItem"
// @Success 200 {object} models.QuoteItem
// @Router /QuoteItems [post]
// @Security ApiKeyAuth
func PostQuoteItem(c *gin.Context) {
	service := services.NewQuoteItemService(daos.NewQuoteItemDAO())
	quoteItem := &models.QuoteItem{}

	if err := c.Bind(quoteItem); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		if result, err := service.Post(quoteItem); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetQuoteItems godoc
// @Summary Retrieves orders
// @Produce json
// @Success 200 [] models.QuoteItem
// @Router /{id}/quoteItems [get]
// @Security ApiKeyAuth
func GetQuoteItems(c *gin.Context) {
	service := services.NewQuoteItemService(daos.NewQuoteItemDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if quoteItems, err := service.GetAll(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, quoteItems)
	}
}

// DeleteQuoteItem godoc
// @Summary Delete quoteItem based on given ID, and order ID
// @Produce json
// @Param id path integer true "Quote ID"
// @Success 204 {object}
// @Router /quoteItems/{itemId} [get]
// @Security ApiKeyAuth
func DeleteQuoteItem(c *gin.Context) {
	service := services.NewQuoteItemService(daos.NewQuoteItemDAO())
	itemId, _ := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err := service.Delete(uint(itemId)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}


// PatchQuoteItem godoc
// @Summary Delete orderItem based on given ID
// @Produce json
// @Param id path integer true "QuoteItem ID"
// @Success 204 {object}
// @Router /quoteItem/{id} [patch]
// @Security ApiKeyAuth
func PatchQuoteItem(c *gin.Context) {
	s := services.NewQuoteItemService(daos.NewQuoteItemDAO())
	id, _ := strconv.Atoi(c.Param("id"))
	quoteItem := &models.QuoteItem{}
	if err := c.Bind(quoteItem);  err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		quoteItem.ID = id

		if result, err := s.Patch(quoteItem); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
