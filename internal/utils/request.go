package utils

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
)

func ParseIDParam(ctx *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid ID parameter")
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
