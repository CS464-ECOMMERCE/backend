package services

import (
	"backend/models"
	"backend/storage"
)

type ItemService struct {
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (i *ItemService) CreateItem(item *models.Item) (*models.Item, error) {
	item, err := storage.StorageInstance.Item.CreateItem(item)
	if err != nil {
		return nil,err
	}
	return item,nil
}

func (i *ItemService) Get(id uint) (*models.Item, error) {
	item, err := storage.StorageInstance.Item.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *ItemService) Update(item *models.Item) error {
	err := storage.StorageInstance.Item.Update(item)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemService) Delete(id uint) error {
	err := storage.StorageInstance.Item.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
