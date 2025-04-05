package controllers

import (
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req models.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Role != nil && *req.Role == models.RoleMerchant {
		if req.BusinessName == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "business_name is required for merchant registration"})
			return
		}
	}

	user, err := services.NewUserService().Register(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.NewUserService().Login(&req)
	if err != nil {
		if err.Error() == "unauthorized" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("===========TOKEN==========", response.Token)
	ctx.SetCookie("token", response.Token, 60*60*24, "/", "", false, false)

	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	// Get user ID from JWT token (assuming you've set it in middleware)
	userID := ctx.GetInt("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.NewUserService().UpdateUser(userID, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

