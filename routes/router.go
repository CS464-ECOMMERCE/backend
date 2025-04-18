package routes

import (
	"backend/configs"
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRoutes() {
	health := controllers.NewHealthController()
	product := controllers.NewProductController()
	user := controllers.NewUserController()
	cart := controllers.NewCartController()
	order := controllers.NewOrderController()
	stripeCon := controllers.NewStripeController()
	router := gin.Default()

	// recover from panics and respond with internal server error
	router.Use(gin.Recovery())

	// add prometheus
	middleware.PrometheusInit()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Use(middleware.TrackMetrics())

	// enabling cors
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	v1.GET("/health", health.HealthCheck)

	// create order & checkout
	v1.POST("/create_order", middleware.CheckSession, product.CreateOrder)
	// webhook for stripe event handler
	v1.POST("/stripe_webhook", order.HandleStripeWebhook)

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
	orderRoute.GET("/:id", order.GetOrder)
	orderRoute.GET("/user/:email", order.GetOrderByEmail)
	orderRoute.GET("/merchant", middleware.CheckAuth, order.GetOrdersByMerchant)
	orderRoute.POST("/update", middleware.CheckAuth, order.UpdateOrderStatus)
	orderRoute.POST("/cancel", middleware.CheckAuth, order.CancelOrder)

	// stripe routes
	stripeRoute := v1.Group("/stripe")
	stripeRoute.GET("/:session_id", stripeCon.GetSession)
	stripeRoute.POST("/cancel/:session_id", stripeCon.CancelSession)

	router.Run(":" + configs.PORT)
}
