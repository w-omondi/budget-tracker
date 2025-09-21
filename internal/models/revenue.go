package models

type Revenue struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `gorm:"type:text" json:"description"`
	Source      string  `json:"source"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateRevenueDto struct {
	Amount      float64 `json:"amount" validate:"required,min=1"`
	Description string  `json:"description" validate:"required"`
	Source      string  `json:"source" validate:"required"`
}

type UpdateRevenueDto struct {
	Amount      float64 `json:"amount" validate:"required,min=1"`
	Description string  `json:"description" validate:"required"`
	Source      string  `json:"source" validate:"required"`
}
