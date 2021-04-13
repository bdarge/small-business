package authentication

import (
  "fmt"
  jwt "github.com/appleboy/gin-jwt/v2"
  "github.com/bdarge/sm-api/cmd/invoice/config"
  "github.com/bdarge/sm-api/cmd/invoice/httputil"
  "github.com/bdarge/sm-api/cmd/invoice/models"
  _ "github.com/bdarge/sm-api/cmd/invoice/apis"
  _ "github.com/bdarge/sm-api/cmd/invoice/httputil"
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
      if v, ok := data.(*models.User); ok {
        return jwt.MapClaims{
          identityKey: v.UserName,
          "id": v.ID,
        }
      } else {
        fmt.Println("error getting claims")
      }
      return jwt.MapClaims{}
    },
    IdentityHandler: func(c *gin.Context) interface{} {
      claims := jwt.ExtractClaims(c)
      userId, _ := strconv.Atoi(fmt.Sprintf("%.f", claims["id"]))
      return &models.User{
        Model: models.Model{ID: userId},
        UserName: claims[identityKey].(string),
      }
    },
    Authenticator: func(c *gin.Context) (interface{}, error) {
      var loginVals models.Login
      if err := c.ShouldBind(&loginVals); err != nil {
        return "", jwt.ErrMissingLoginValues
      }
      username := loginVals.Username
      password := loginVals.Password

      var user models.User
      if err := config.Config.DB.Where("userName = ?", username).First(&user).Error; err != nil {
        fmt.Print(err.Error())
        return nil, jwt.ErrFailedAuthentication
      }
      fmt.Printf("userID %d, email %s", user.ID, user.Email)
      match := httputil.CheckPasswordHash(password, user.Password)

      fmt.Println(match)
      if !match {
        return nil, jwt.ErrFailedAuthentication
      }
      fmt.Println("user authenticated")
      return &user, nil
    },
    Authorizator: func(data interface{}, c *gin.Context) bool {
      //if v, ok := data.(*models.User); ok && v.UserName == "admin" {
      //  return true
      //}
      //
      //return false
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
