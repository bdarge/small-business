package authentication

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	_ "github.com/bdarge/api/cmd/sm/apis"
	"github.com/bdarge/api/cmd/sm/config"
	"github.com/bdarge/api/cmd/sm/helper"
	"github.com/bdarge/api/cmd/sm/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Middleware(identityKey string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Account); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"id":        v.ID,
					"userId": 	 v.UserID,
				}
			} else {
				fmt.Println("error getting claims")
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userId, _ := strconv.Atoi(fmt.Sprintf("%.f", claims["id"]))
			return &models.Account{
				Model:    models.Model{ID: userId},
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login models.Login
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := login.Username
			password := login.Password

			var account models.Account
			if err := config.Config.DB.Where("email = ?", username).First(&account).Error; err != nil {
				fmt.Print(err.Error())
				return nil, jwt.ErrFailedAuthentication
			}
			fmt.Printf("userID %d, email %s", account.ID, account.Email)
			match := helper.CheckPasswordHash(password, account.Password)

			fmt.Println(match)
			if !match {
				return nil, jwt.ErrFailedAuthentication
			}
			fmt.Println("account authenticated")
			return &account, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// for this sample exercise all users are authorized
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value.
		// This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
