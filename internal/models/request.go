package models

type QueryOptions struct {
	Query    string `json:"query" validate:"required"`
	Page     int    `json:"page" validate:"required,min=1"`
	PageSize int    `json:"page_size" validate:"required,min=1"`
}
