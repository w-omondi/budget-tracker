package utils

import (
	"fmt"

	"github.com/w-omondi/budget-tracker.git/internal/models"
)

type ApiResponse[T any] interface {
	SendSuccessResponse(message string, data any) *models.NormalResponse
	SendPaginatedResponse(page, pageSize, totalItems int, items []T) *models.PaginatedResponse[T]
	SendErrorResponse(message string, err error) *models.ErrorResponse
}

type apiResponse[T any] struct{}

func NewApiResponse[T any]() ApiResponse[T] {
	return &apiResponse[T]{}
}

func (r *apiResponse[T]) SendSuccessResponse(message string, data any) *models.NormalResponse {
	return &models.NormalResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func (r *apiResponse[T]) SendPaginatedResponse(page, pageSize, totalItems int, items []T) *models.PaginatedResponse[T] {
	fmt.Printf("Response size: Page: %v, PageSize: %v, Data Sent: %v, TotalItems: %v", page, pageSize, len(items), totalItems)
	return &models.PaginatedResponse[T]{
		Success:    true,
		Message:    "Data retrieved successfully",
		Data:       items,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: (totalItems + pageSize - 1) / pageSize,
	}
}

func (r *apiResponse[T]) SendErrorResponse(message string, err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Success: false,
		Message: message,
		Error:   err.Error(),
	}
}
