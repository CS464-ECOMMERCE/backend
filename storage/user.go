package storage

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserStorage struct {
	read  *gorm.DB
	write *gorm.DB
}

type UserInterface interface {
	Create(user *models.User) error
	CreateMerchant(merchant *models.Merchant) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id int) (*models.User, error)
	Update(user *models.User, updates map[string]interface{}) error
	UpdateMerchant(userID int, updates map[string]interface{}) error
}

func NewUserTable(read, write *gorm.DB) UserInterface {
	StorageInstance.AutoMigrate(&models.User{})
	return &UserStorage{
		read:  read,
		write: write,
	}
}

func (s *UserStorage) Create(user *models.User) error {
	return s.write.Create(user).Error
}

func (s *UserStorage) CreateMerchant(merchant *models.Merchant) error {
	return s.write.Create(merchant).Error
}

func (s *UserStorage) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.read.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	
	return &user, nil
}

func (s *UserStorage) FindByID(id int) (*models.User, error) {
	var user models.User
	if err := s.read.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStorage) Update(user *models.User, updates map[string]interface{}) error {
	return s.write.Model(user).Updates(updates).Error
}

func (s *UserStorage) UpdateMerchant(userID int, updates map[string]interface{}) error {
	return s.write.Model(&models.Merchant{}).Where("user_id = ?", userID).Updates(updates).Error
}
