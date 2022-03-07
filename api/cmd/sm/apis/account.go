package apis

import (
	"errors"
	"fmt"
	"github.com/bdarge/api/cmd/sm/daos"
	"github.com/bdarge/api/cmd/sm/models"
	"github.com/bdarge/api/cmd/sm/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// GetAccount godoc
// @Summary Retrieves account based on given ID
// @Produce json
// @Param id path integer true "Account ID"
// @Success 200 {object} models.Account
// @Router /accounts/{id} [get]
// @Security ApiKeyAuth
func GetAccount(c *gin.Context) {
	s := services.NewAccountService(daos.NewAccountDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if acct, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, acct)
	}
}

// GetAccountUser godoc
// @Summary Retrieves account based on given ID
// @Produce json
// @Param id path integer true "Account ID"
// @Success 200 {object} models.User
// @Router /accounts/{id}/user [get]
// @Security ApiKeyAuth
func GetAccountUser(c *gin.Context) {
	u := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := u.GetByAccountId(uint(id)); err != nil {
		fmt.Printf("user not found by account id. %v", id)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}


// PostAccount godoc
// @Summary Creates an account
// @Param user body models.Account true "Add account"
// @Success 200 {object} models.Account
// @Router /Account [post]
// @Security ApiKeyAuth
func PostAccount(c *gin.Context) {
	register := &models.Register{}
	if err := c.Bind(register);  err != nil {
		fmt.Printf("invalid post data. %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		a := services.NewAccountService(daos.NewAccountDAO())
		if !isEmailValid(register.Email) {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Printf("invalid email. %v Email=%v", err, register.Email)
			return
		}

		if acct, error:= a.GetByEmail(register.Email); error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			if acct != nil {
				c.AbortWithError(http.StatusBadRequest, errors.New("account exists"))
				return
			}
		}

		acct := &models.Account{
			Password: register.Password,
			Email:    register.Email,
		}

		if result, err := a.Post(acct); err != nil {
			fmt.Printf("user creation has failed: %v", err)
			fmt.Println(result)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

