package models

type UserRole string

const (
	RoleClient   UserRole = "client"
	RoleMerchant UserRole = "merchant"
	RoleAdmin    UserRole = "admin"
)

type User struct {
	ID           int      `json:"id"`
	Email        string   `json:"email" gorm:"unique"`
	PasswordHash string   `json:"-"` // "-" means this won't be included in JSON
	Role         UserRole `json:"role"`
}

type Merchant struct {
	UserID       int    `json:"user_id"`
	BusinessName string `json:"business_name"`
	TaxID        string `json:"tax_id,omitempty"`
	Verified     bool   `json:"verified"`
}

// Request/Response structs
type RegisterUserRequest struct {
	Email        string    `json:"email" binding:"required,email"`
	Password     string    `json:"password" binding:"required_if=Role merchant,min=6"`
	Role         *UserRole `json:"role,omitempty"`
	BusinessName string    `json:"business_name,omitempty" binding:"required_if=Role merchant"`
	TaxID        string    `json:"tax_id,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	BusinessName string `json:"business_name,omitempty"` // For merchants only
	TaxID        string `json:"tax_id,omitempty"`        // For merchants only
}
