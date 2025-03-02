package models

type Item struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *uint  `json:"deleted_at" gorm:"autoDeleteTime"`
}

