package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/w-omondi/budget-tracker.git/internal/models"
)

func ParseIDParam(c *fiber.Ctx) (uuid.UUID, error) {
	idParam := c.Params("id")
	if idParam == "" {
		return uuid.Nil, errors.New("missing ID parameter")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID format")
	}

	return id, nil
}

func ParseQueryOptions(ctx *fiber.Ctx) *models.QueryOptions {
	page := ctx.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	pageSize := ctx.QueryInt("page_size", 10)
	if pageSize < 1 {
		pageSize = 10
	}

	query := ctx.Query("query", "")

	if query != "" {
		return &models.QueryOptions{
			Query:    query,
			Page:     page,
			PageSize: pageSize,
		}
	}

	return &models.QueryOptions{
		Page:     page,
		PageSize: pageSize,
	}
}
