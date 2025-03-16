package routes

import (
	"backend/configs"
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// swaggerFiles "github.com/swaggo/files"     // swagger embed files
	// ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func InitRoutes() {

	// challenge := controllers.NewChallengeController(configs.Client)
	// image := controllers.NewImageController(configs.Client)
	// process := controllers.NewProcessController(configs.Client)
	// attempt := controllers.NewAttemptController(configs.Client)

	health := controllers.NewHealthController()
	item := controllers.NewItemController()
	product := controllers.NewProductController()
	user := controllers.NewUserController()
	router := gin.Default()

	// // recover from panics and respond with internal server error
	router.Use(gin.Recovery())

	// add prometheus
	middleware.PrometheusInit()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Use(middleware.TrackMetrics())

	// enabling cors
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")

	v1.GET("/health", health.HealthCheck)

	itemRoute := v1.Group("/item")
	itemRoute.POST("", item.CreateItem)
	itemRoute.GET("/:id", item.GetItem)
	itemRoute.PUT("", item.UpdateItem)
	itemRoute.DELETE("/:id", item.DeleteItem)

	productRoute := v1.Group("/product")
	productRoute.GET("", product.GetProduct)
	productRoute.POST("", product.CreateProduct)
	productRoute.PATCH("", product.UpdateProduct)
	productRoute.DELETE("/:id", product.DeleteProduct)
	productRoute.GET("/:id", product.GetProductById)

	userRoute := v1.Group("/user")
	userRoute.POST("/register", user.Register)
	userRoute.POST("/login", user.Login)

	// // Swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + configs.PORT)
}
