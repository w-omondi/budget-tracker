package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Revenue struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Amount      float64 `json:"amount"`
	Description string  `gorm:"type:text" json:"description"`
	Source      string  `json:"source"`
}

type CreateRevenueDto struct {
	Amount      float64 `json:"amount" validate:"required,min=1"`
	Description string  `json:"description" validate:""`
	Source      string  `json:"source" validate:"required"`
}

type UpdateRevenueDto struct {
	Amount      float64 `json:"amount" validate:"required,min=1"`
	Description string  `json:"description" validate:""`
	Source      string  `json:"source" validate:"required"`
}
