package controllers

import (
	"backend/services"
	"net/http"
	pb "backend/proto"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (p *ProductController) GetProduct(c *gin.Context) {
	products, err := services.NewProductService().GetProduct(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product pb.CreateProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := services.NewProductService().CreateProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

