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

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// PostUser godoc
// @Summary Creates a user
// @Param user body models.User true "Add user"
// @Success 200 {object} models.User
// @Router /users [post]
// @Security ApiKeyAuth
func PostUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	user := &models.User{}
	if err := c.Bind(user);  err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		if user, err := s.Post(user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// PatchUser godoc
// @Summary Creates a user
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Router /users/{id} [patch]
// @Security ApiKeyAuth
func PatchUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		if err := c.Bind(user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Println(err)
		} else {
			if user, err := s.Patch(user); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				log.Println(err)
			} else {
				c.JSON(http.StatusOK, user)
			}
		}
	}
}