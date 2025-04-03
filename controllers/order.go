package controllers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (o *OrderController) GetOrder(c *gin.Context) {
	var req struct {
		Id uint64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := services.NewOrderService().GetOrder(c, req.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

func (o *OrderController) GetOrderByEmail(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "email is required"})
		return
	}

	user, err := services.NewUserService().GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	orders, err := services.NewOrderService().GetOrdersByUser(c, uint64(user.ID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}

func (o *OrderController) GetOrdersByMerchant(c *gin.Context) {
	// Get user ID from JWT token (assuming you've set it in middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := services.NewOrderService().GetOrdersByMerchant(c, uint64(userID.(float64)))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}

func (o *OrderController) UpdateOrderStatus(c *gin.Context) {
	var req struct {
		Id     uint64 `json:"id"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := services.NewOrderService().UpdateOrderStatus(c, req.Id, req.Status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

func (o *OrderController) CancelOrder(c *gin.Context) {
	var req struct {
		Id uint64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := services.NewOrderService().CancelOrder(c, req.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	var req struct {
		Id uint64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := services.NewOrderService().DeleteOrder(c, req.Id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
