package utils

import (
	"log"

	"github.com/w-omondi/budget-tracker.git/internal/models"
)

type ApiResponse[T any] interface {
	SendNoContedResponse() *models.NoContentResponse
	SendSingleResponse(data any) *models.NormalResponse
	SendPaginatedResponse(page, pageSize, totalItems int, items []T) *models.PaginatedResponse[T]
	SendErrorResponse(message string, err error) *models.ErrorResponse
}

type apiResponse[T any] struct{}

func NewApiResponse[T any]() ApiResponse[T] {
	return &apiResponse[T]{}
}

func (r *apiResponse[T]) SendNoContedResponse() *models.NoContentResponse {
	return &models.NoContentResponse{
		Success: true,
		Message: "No content",
	}
}

func (r *apiResponse[T]) SendSingleResponse(data any) *models.NormalResponse {
	return &models.NormalResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    data,
	}
}

func (r *apiResponse[T]) SendPaginatedResponse(page, pageSize, totalItems int, items []T) *models.PaginatedResponse[T] {
	log.Printf("Response size: Page: %v, PageSize: %v, Data Sent: %v, TotalItems: %v", page, pageSize, len(items), totalItems)
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
	log.Printf("Error: %s", message)
	if err != nil {
		log.Printf("Details: %v", err.Error())
	}
	return &models.ErrorResponse{
		Success: false,
		Message: message,
	}
}
