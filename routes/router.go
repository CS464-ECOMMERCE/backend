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
	health := controllers.NewHealthController()
	product := controllers.NewProductController()
	user := controllers.NewUserController()
	cart := controllers.NewCartController()
	order := controllers.NewOrderController()
	stripeCon := controllers.NewStripeController()
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
	// config.AllowOrigins = []string{
	// 	"http://localhost:3000",
	// 	"https://*.stripe.com",
	// }
	// config.AllowCredentials = true
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	v1.GET("/health", health.HealthCheck)

	v1.POST("/create_order", middleware.CheckSession, product.CreateOrder)

	// product routes
	productRoute := v1.Group("/product")
	productRoute.GET("", product.GetProduct)
	productRoute.GET("/merchant", middleware.CheckAuth, product.GetProductByMerchantId)
	productRoute.POST("", middleware.CheckAuth, product.CreateProduct)
	productRoute.PATCH("", middleware.CheckAuth, product.UpdateProduct)
	productRoute.DELETE("/:id", middleware.CheckAuth, product.DeleteProduct)
	productRoute.GET("/:id", product.GetProductById)
	productRoute.POST("/upload/:id", middleware.CheckAuth, product.UpdateProductImages)

	// user routes
	userRoute := v1.Group("/user")
	userRoute.POST("/register", user.Register)
	userRoute.POST("/login", user.Login)

	// cart routes
	cartRoute := v1.Group("/cart")
	cartRoute.POST("/add", middleware.CheckSession, cart.AddItem)
	cartRoute.GET("", middleware.CheckSession, cart.GetCart)
	cartRoute.POST("/update", middleware.CheckSession, cart.UpdateItemQuantity)
	cartRoute.POST("/remove_item", middleware.CheckSession, cart.RemoveItem)
	cartRoute.POST("/empty", middleware.CheckSession, cart.EmptyCart)

	// order routes
	orderRoute := v1.Group("/order")
	orderRoute.GET("", order.GetOrder)
	orderRoute.GET("/user", order.GetOrderByEmail)
	orderRoute.GET("/merchant", middleware.CheckAuth, order.GetOrdersByMerchant)
	orderRoute.POST("/update", middleware.CheckAuth, order.UpdateOrderStatus)
	orderRoute.POST("/cancel", middleware.CheckAuth, order.CancelOrder)
	orderRoute.POST("/delete", middleware.CheckAuth, order.DeleteOrder)

	// stripe webhook routes
	stripeRoute := v1.Group("/stripe")
	stripeRoute.POST("/webhook", middleware.CheckSession, stripeCon.Webhook)

	router.Run(":" + configs.PORT)
}
