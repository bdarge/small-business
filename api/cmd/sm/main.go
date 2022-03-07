package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/bdarge/api/cmd/sm/apis"
	"github.com/bdarge/api/cmd/sm/authentication"
	. "github.com/bdarge/api/cmd/sm/config"
	"github.com/bdarge/api/cmd/sm/helper"
	"github.com/bdarge/api/cmd/sm/migrator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/bdarge/api/cmd/sm/docs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

// @title SM Swagger API
// @version 1.0
// @description Swagger API for Business X.
// @termsOfService http://swagger.io/terms/

// @BasePath /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	filename := os.Getenv("PATH_TO_CONFIG")
	// load application configurations
	if err := LoadConfig(filename, "./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s. filename: %s", err, filename))
	}

	// init env object
	env := helper.NewEnv()
	migrateData := env.GetBool("MIGRATOR")
	log.Printf("migrateData => %t", migrateData)

	if migrateData {
		if _, err := migrator.Migrate(); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successfully migrated database.")
		os.Exit(0)
	}

	// Creates a router without any middleware by default
	router := gin.New()

	//Default() allows all origins
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/swagger", func(c *gin.Context) {
		// fmt.Print(c.FullPath())
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	url := ginSwagger.URL(Config.BaseUrl + "/swagger/doc.json")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	var identityKey = "email"

	// the jwt middleware
	authMiddleware, err := authentication.Middleware(identityKey)

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ok"})
	})

	auth := router.Group("/v1/auth")
	{
		// Refresh time can be longer than token timeout
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.POST("/login", authMiddleware.LoginHandler)
	}

	v1 := router.Group("/v1")
	{
		v1.POST("/accounts", apis.PostAccount)
		v1.Use(authMiddleware.MiddlewareFunc())
		v1.GET("/accounts/:id", apis.GetAccount)
		v1.GET("/accounts/:id/user", apis.GetAccountUser)

		// user
		v1.GET("/users/:id", apis.GetUser)
		v1.POST("/users", apis.PostUser)
		v1.PATCH("/users/:id", apis.PatchUser)

		// customer
		v1.GET("/customers", apis.GetCustomers)
		v1.GET("/customers/:id", apis.GetCustomer)
		v1.POST("/customers", apis.PostCustomer)
		v1.DELETE("/customers/:id", apis.DeleteCustomer)
		v1.PATCH("/customers/:id", apis.PatchCustomer)

		// orders
		v1.GET("/orders", apis.GetOrders)
		v1.GET("/orders/:id", apis.GetOrder)
		v1.POST("/orders", apis.PostOrder)
		v1.DELETE("/orders/:id", apis.DeleteOrder)
		v1.PATCH("/orders/:id", apis.PatchOrder)

		// quote
		v1.GET("/quotes", apis.GetQuotes)
		v1.GET("/quotes/:id", apis.GetQuote)
		v1.POST("/quotes", apis.PostQuote)
		v1.DELETE("/quotes/:id", apis.DeleteQuote)
		v1.PATCH("/quotes/:id", apis.PatchQuote)

		// order item
		v1.GET("/orders/:id/items", apis.GetOrderItems)
		v1.POST("/orders/:id/items", apis.PostOrderItem)
		v1.DELETE("/order-items/:itemId", apis.DeleteOrderItem)
		v1.PATCH("/order-items/:itemId", apis.PatchOrderItem)

		//quote item
		v1.GET("/quotes/:id/items", apis.GetQuoteItems)
		v1.POST("/quotes/:id/items", apis.PostQuoteItem)
		v1.DELETE("/quote-items/:itemId", apis.DeleteQuoteItem)
		v1.PATCH("/quote-items/:itemId", apis.PatchQuoteItem)
	}

	log.Printf("Open database, %s", Config.Database)
	// open database
	Config.DB, Config.DBErr = gorm.Open(mysql.Open(Config.DSN), &gorm.Config{})

	if Config.DBErr != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := Config.SqlDB.Close(); err != nil {
			log.Fatal("failed to close db:" + err.Error())
		}
	}()

	log.Println("Successfully connected to database")

	err = router.Run(fmt.Sprintf(":%v", Config.ServerPort))
	if err != nil {
		log.Fatal("gin router Error:" + err.Error())
	}
}
