package controllers

import (
	pb "backend/proto"
	"backend/services"
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (p *ProductController) GetProduct(c *gin.Context) {
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	cursor, err := strconv.ParseUint(c.DefaultQuery("cursor", "0"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor"})
		return
	}

	products, err := services.NewProductService().GetProduct(c.Request.Context(), &pb.ListProductsRequest{
		Cursor: cursor,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductByMerchantId(c *gin.Context) {
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	cursor, err := strconv.ParseUint(c.DefaultQuery("cursor", "0"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor"})
		return
	}

	merchantId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	products, err := services.NewProductService().GetProductByMerchantId(c.Request.Context(), &pb.ListProductsRequest{
		MerchantId: uint64(merchantId.(float64)),
		Cursor:     cursor,
		Limit:      limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product pb.CreateProductRequest
	merchantId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}
	product.MerchantId = uint64(merchantId.(float64))
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

func (p *ProductController) UpdateProduct(c *gin.Context) {
	var product pb.UpdateProductRequest
	merchantId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}
	product.MerchantId = uint64(merchantId.(float64))

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := services.NewProductService().UpdateProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	var product pb.DeleteProductRequest
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product.Id = id

	merchantId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}
	product.MerchantId = uint64(merchantId.(float64))

	_, err = services.NewProductService().DeleteProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (p *ProductController) GetProductById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := services.NewProductService().GetProductById(c.Request.Context(), &pb.GetProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (p *ProductController) UpdateProductImages(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}
	files := c.Request.MultipartForm.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images uploaded"})
		return
	}
	for _, file := range files {
		contentType := file.Header.Get("Content-Type")
		if contentType != "image/jpeg" && contentType != "image/png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPEG and PNG images are allowed"})
			log.Printf("Invalid content type: %s", contentType)
			return
		}
		if file.Size > 10<<20 { // 10 MB
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB"})
			return
		}
	}

	resp, err := services.NewProductService().UpdateProductImages(c.Request.Context(), *c.Request.MultipartForm, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false) // Prevent escaping of `&`
	encoder.Encode(resp)
	return
}

func (p *ProductController) CreateOrder(c *gin.Context) {
	sessionID, exists := c.Get("session_id")
	if !exists {
		c.JSON(400, gin.H{"error": "Invalid session ID"})
		return
	}

	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "email is required"})
		return
	}

	user, err := services.NewUserService().CreateBuyerAccountIfNotExist(req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	placeOrderRequest := &pb.PlaceOrderRequest{
		SessionId: sessionID.(string),
		UserEmail: req.Email,
		UserId:    uint64(user.ID),
	}

	sess, err := services.NewProductService().PlaceOrder(c, placeOrderRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"checkoutUrl": sess.CheckoutUrl})
}
