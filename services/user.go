package services

import (
	"backend/configs"
	"backend/models"
	"backend/storage"
	"errors"

	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	jwtSecret []byte
}

func NewUserService() *UserService {
	return &UserService{
		jwtSecret: []byte(configs.JWT_SECRET),
	}
}

func (s *UserService) Register(req *models.RegisterUserRequest) (*models.User, error) {
	// Check if email already exists\
	existingUser, err := storage.StorageInstance.User.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Set default role if not provided
	role := models.RoleClient
	if req.Role != nil {
		role = *req.Role
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         role,
	}

	if err := storage.StorageInstance.User.Create(user); err != nil {
		return nil, err
	}

	// If registering as merchant, create merchant record
	if role == models.RoleMerchant {
		merchant := &models.Merchant{
			UserID:       user.ID,
			BusinessName: req.BusinessName,
			TaxID:        req.TaxID,
			Verified:     false,
		}

		if err := storage.StorageInstance.User.CreateMerchant(merchant); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *UserService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	user, err := storage.StorageInstance.User.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateJWT(*user)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}

func (s *UserService) UpdateUser(userID int, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := storage.StorageInstance.User.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	updates := map[string]interface{}{}

	if req.Email != "" {
		updates["email"] = req.Email
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updates["password_hash"] = string(hashedPassword)
	}

	if err := storage.StorageInstance.User.Update(user, updates); err != nil {
		return nil, err
	}

	// Update merchant details if applicable
	if user.Role == models.RoleMerchant && (req.BusinessName != "" || req.TaxID != "") {
		merchantUpdates := map[string]interface{}{}
		if req.BusinessName != "" {
			merchantUpdates["business_name"] = req.BusinessName
		}
		if req.TaxID != "" {
			merchantUpdates["tax_id"] = req.TaxID
		}

		if err := storage.StorageInstance.User.UpdateMerchant(user.ID, merchantUpdates); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *UserService) generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString(s.jwtSecret)
}
