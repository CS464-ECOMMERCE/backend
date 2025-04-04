package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

type StripeController struct{}

func NewStripeController() *StripeController {
	return &StripeController{}
}

func (s *StripeController) CancelSession(c *gin.Context) {
	sessionId := c.Param("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session"})
		return
	}

	params := &stripe.CheckoutSessionExpireParams{}
	result, err := session.Expire(sessionId, params)
	if err != nil || result.Status != stripe.CheckoutSessionStatusExpired {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to expire session: %v", err)})
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

func (s *StripeController) GetSession(c *gin.Context) {
	sessionId := c.Param("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session"})
		return
	}

	params := &stripe.CheckoutSessionParams{}
	result, err := session.Get(sessionId, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to get session: %v", err)})
		return
	}

	c.JSON(http.StatusOK, result)
}
