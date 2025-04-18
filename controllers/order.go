package controllers

import (
	"backend/configs"
	"backend/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (o *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "order id is required"})
		return
	}

	orderId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid order id"})
		return
	}

	order, err := services.NewOrderService().GetOrder(c, orderId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

func (o *OrderController) GetOrderByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(400, gin.H{"error": "email is required"})
		return
	}

	user, err := services.NewUserService().GetUserByEmail(email)
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

func (o *OrderController) HandleStripeWebhook(c *gin.Context) {
	event := o.verifyEvent(c)

	if event == nil {
		c.JSON(http.StatusOK, "Ok")
		return
	}

	eventsToListen := []stripe.EventType{
		stripe.EventTypeCheckoutSessionCompleted,
		stripe.EventTypeCheckoutSessionExpired,
	}

	var session stripe.CheckoutSession
	if slices.Contains(eventsToListen, stripe.EventType(event.Type)) {
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			c.JSON(400, gin.H{"error": fmt.Sprintf("unable to unmarshal data: %v", err.Error())})
			return
		}
	} else {
		c.JSON(200, "Ok")
		return
	}

	orderId, err := strconv.ParseUint(session.Metadata["orderId"], 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("unable to get orderId: %v", err.Error())})
		return
	}

	if err := services.NewOrderService().UpdatePaymentStatus(c, fmt.Sprint(event.Type), orderId); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "success"})
}

func (o *OrderController) verifyEvent(c *gin.Context) *stripe.Event {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Error reading request body: %v", err)})
		return nil
	}

	var event stripe.Event

	if err := json.Unmarshal(payload, &event); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Error binding JSON: %v", err)})
		return nil
	}

	signatureHeader := c.GetHeader("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, configs.STRIPE_API_KEY)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Error verifying webhook signature: %v", err)})
		return nil
	}

	return &event
}
