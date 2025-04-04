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

func (s *StripeController) getCheckoutSession(sessionId string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{}
	result, err := session.Get(sessionId, params)
	return result, err
}

func (s *StripeController) CancelSession(c *gin.Context) {
	sessionId := c.Param("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session"})
		return
	}

	res, err := s.getCheckoutSession(sessionId)
	if err != nil {
		// Even here, stripe may throw 400 if session not found
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to get session: %v", err)})
		return
	}

	if res.Status != stripe.CheckoutSessionStatusOpen {
		c.JSON(http.StatusOK, gin.H{"message": "session is already not open"})
		return
	}

	_, err = session.Expire(sessionId, &stripe.CheckoutSessionExpireParams{})
	if err != nil {
		// Still return 200 to prevent Gin from trying to write twice
		c.JSON(http.StatusOK, gin.H{"message": "could not expire session, possibly already expired"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "session expired successfully"})
}

func (s *StripeController) GetSession(c *gin.Context) {
	sessionId := c.Param("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session"})
		return
	}

	result, err := s.getCheckoutSession(sessionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to get session: %v", err)})
		return
	}

	c.JSON(http.StatusOK, result)
}
