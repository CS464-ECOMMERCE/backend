package controllers

import (
	"backend/models"
	"backend/services"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
}

func NewItemController() *ItemController {
	return &ItemController{}
}

func (i *ItemController) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdItem, err := services.NewItemService().CreateItem(&item)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createdItem)
}

func (i *ItemController) GetItem(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.ParseUint(itemID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	item, err := services.NewItemService().Get(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, item)
}

func (i *ItemController) UpdateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := services.NewItemService().Update(&item)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, item)
}

func (i *ItemController) DeleteItem(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.ParseUint(itemID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	err = services.NewItemService().Delete(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}
