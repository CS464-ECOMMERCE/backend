package storage

import (
	"backend/models"

	"gorm.io/gorm"
)

type ItemInterface interface {
	CreateItem(item *models.Item) (*models.Item, error)
	Get(id uint) (*models.Item, error)
	Update(item *models.Item) error
	Delete(id uint) error
}

type itemDB struct {
	read  *gorm.DB
	write *gorm.DB
}

// CreateItem implements ItemInterface.
func (i *itemDB) CreateItem(item *models.Item) (*models.Item, error) {
	ret := i.write.Create(item)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return item, nil
}

// Delete implements ItemInterface.
func (i *itemDB) Delete(id uint) error {
	return i.write.Delete(&models.Item{}, id).Error
}

// Get implements ItemInterface.
func (i *itemDB) Get(id uint) (*models.Item, error) {
	item := &models.Item{}
	ret := i.read.First(item, id)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return item, nil
}

// Update implements ItemInterface.
func (i *itemDB) Update(item *models.Item) error {
	return i.write.Save(item).Error
}

func NewItemTable(read, write *gorm.DB) ItemInterface {
	StorageInstance.AutoMigrate(&models.Item{})
	return &itemDB{
		read:  read,
		write: write,
	}
}
