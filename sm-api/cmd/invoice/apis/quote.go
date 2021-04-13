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

// QuoteDAO persists user data in database
type QuoteDAO struct{}

// NewQuoteDAO creates a new NewQuoteDAO
func NewQuoteDAO() *QuoteDAO {
	return &QuoteDAO{}
}

// GetQuote godoc
// @Summary Retrieves quote based on given ID
// @Produce json
// @Param id path integer true "Quote ID"
// @Success 200 {object} models.Quote
// @Router /quotes/{id} [get]
// @Security ApiKeyAuth
func GetQuote(c *gin.Context) {
	s := services.NewQuoteService(daos.NewQuoteDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// PostQuote godoc
// @Summary Creates a quote
// @Param user body models.Register true "Add quote"
// @Success 200 {object} models.Quote
// @Router /Quotes [post]
// @Security ApiKeyAuth
func PostQuote(c *gin.Context) {
	service := services.NewQuoteService(daos.NewQuoteDAO())
	quote := &models.Quote{}

	if err := c.Bind(quote); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		if result, err := service.Post(quote); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetQuotes godoc
// @Summary Retrieves quotes
// @Produce json
// @Success 200 [] models.Quote
// @Router /quotes [get]
// @Security ApiKeyAuth
func GetQuotes(c *gin.Context) {
	service := services.NewQuoteService(daos.NewQuoteDAO())
	var query = models.Query{}
	if err := c.BindQuery(&query); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("query %s", query)
		if quotePage, err := service.GetAll(query); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, quotePage)
		}
	}
}

// DeleteQuote godoc
// @Summary Delete quote based on given ID
// @Produce json
// @Param id path integer true "Quote ID"
// @Success 204 {object}
// @Router /quotes/{id} [get]
// @Security ApiKeyAuth
func DeleteQuote(c *gin.Context) {
	service := services.NewQuoteService(daos.NewQuoteDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := service.Delete(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}


// PatchQuote godoc
// @Summary Delete order based on given ID
// @Produce json
// @Param id path integer true "Quote ID"
// @Success 204 {object}
// @Router /quotes/{id} [get]
// @Security ApiKeyAuth
func PatchQuote(c *gin.Context) {
	s := services.NewQuoteService(daos.NewQuoteDAO())
	id, _ := strconv.Atoi(c.Param("id"))
	quote := &models.Quote{}
	if err := c.Bind(quote);  err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		quote.ID = id
		if result, err := s.Patch(quote); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}