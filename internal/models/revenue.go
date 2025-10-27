package models

import (
	"time"

	"gorm.io/gorm"
)

type Revenue struct {
	// *gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Source      string  `json:"source,omitempty"`
}

type CreateRevenueDto struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Source      string  `json:"source,omitempty"`
}
