package models

type NormalResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type NoContentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PaginatedResponse[T any] struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       []T    `json:"data"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	TotalItems int    `json:"total_items"`
	TotalPages int    `json:"total_pages"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
