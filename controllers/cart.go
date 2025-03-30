package controllers

import (
	pb "backend/proto"
	"backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct{}

func NewCartController() *CartController {
	return &CartController{}
}

func (c *CartController) AddItem(ctx *gin.Context) {

	sessionID, exists := ctx.Get("session_id")
	if !exists {
		sessionID = uuid.New().String()
		ctx.SetCookie("session_id", sessionID.(string), 60*60*24, "/", "", false, false)
	}

	var item pb.CartItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := services.NewCartService().AddItem(ctx, sessionID.(string), &item); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}

func (c *CartController) GetCart(ctx *gin.Context) {
	sessionID, exists := ctx.Get("session_id")
	if !exists {
		ctx.JSON(200, &pb.Cart{})
		return
	}

	cart, err := services.NewCartService().GetCart(ctx, sessionID.(string))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, cart)
}

func (c *CartController) EmptyCart(ctx *gin.Context) {
	sessionID, exists := ctx.Get("session_id")
	if !exists {
		ctx.JSON(200, gin.H{"status": "success"})
		return
	}

	if err := services.NewCartService().EmptyCart(ctx, sessionID.(string)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}

func (c *CartController) RemoveItem(ctx *gin.Context) {
	sessionID, exists := ctx.Get("session_id")
	if !exists {
		ctx.JSON(200, gin.H{"status": "success"})
		return
	}

	var req struct {
		Id uint64 `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := services.NewCartService().RemoveItem(ctx, sessionID.(string), req.Id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}

func (c *CartController) UpdateItemQuantity(ctx *gin.Context) {
	sessionID, exists := ctx.Get("session_id")
	if !exists {
		ctx.JSON(200, gin.H{"status": "success"})
		return
	}

	var item pb.CartItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := services.NewCartService().UpdateItemQuantity(ctx, sessionID.(string), item.Id, item.Quantity); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}
