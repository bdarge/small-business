package apis

import (
  "github.com/bdarge/sm-api/cmd/invoice/daos"
  "github.com/bdarge/sm-api/cmd/invoice/models"
  "github.com/bdarge/sm-api/cmd/invoice/services"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "regexp"
  "strconv"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

// PostUser godoc
// @Summary Creates a user
// @Param user body models.Register true "Add user"
// @Success 200 {object} models.User
// @Router /Register [post]
// @Security ApiKeyAuth
func RegisterUser(c *gin.Context) {
  s := services.NewUserService(daos.NewUserDAO())
  register := &models.Register{}
  if err := c.Bind(register);  err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
    log.Println(err)
  } else {
    user := &models.User{
      Password: register.Password,
      UserName: register.Username,
    }

    if isEmailValid(register.Username) {
      user.Email = register.Username
    }

    if user, err := s.Post(user); err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
      log.Println(err)
    } else {
      c.JSON(http.StatusOK, user)
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
