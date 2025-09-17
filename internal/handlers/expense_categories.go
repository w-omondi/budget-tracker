package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"github.com/w-omondi/budget-tracker.git/internal/utils"
)

type ExpenseCategoriesHandler interface {
	CreateExpenseCategoriesHandler(ctx *fiber.Ctx) error
	GetExpenseCategoryByIDHandler(ctx *fiber.Ctx) error
	GetAllExpenseCategoriesHandler(ctx *fiber.Ctx) error
	UpdateExpenseCategoriesHandler(ctx *fiber.Ctx) error
	DeleteExpenseCategoriesHandler(ctx *fiber.Ctx) error
}

type expenseCategoriesHandler struct {
	expenseCategoriesService services.ExpenseCategoryService
	responseUtil             utils.ApiResponse[*models.ExpenseCategory]
}

func NewExpenseCategoriesHandler(service services.ExpenseCategoryService) ExpenseCategoriesHandler {
	return &expenseCategoriesHandler{
		expenseCategoriesService: service,
		responseUtil:             utils.NewApiResponse[*models.ExpenseCategory](),
	}
}

func (h *expenseCategoriesHandler) CreateExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	log.Println("Handling creation of expenseCategory:")

	expenseCategory := new(models.CreateExpenseCategoryDto)
	if err := utils.ParseAndValidateData(ctx, expenseCategory); err != nil {
		response := h.responseUtil.SendErrorResponse(err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.expenseCategoriesService.CreateCategory(expenseCategory); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to create expenseCategory", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) GetExpenseCategoryByIDHandler(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid expenseCategory ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	expenseCategory, err := h.expenseCategoriesService.GetCategoryByID(id)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to fetch expenseCategory", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if expenseCategory == nil {
		response := h.responseUtil.SendErrorResponse("ExpenseCategory not found", nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}

	response := h.responseUtil.SendSingleResponse(expenseCategory)
	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) GetAllExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	log.Println("Fetching all expenseCategories")

	queryOptions := &models.QueryOptions{
		Query:    ctx.Query("query", ""),
		Page:     ctx.QueryInt("page", 1),
		PageSize: ctx.QueryInt("page_size", 10),
	}

	expenseCategories, count, err := h.expenseCategoriesService.GetAllCategories(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to fetch expenseCategories", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendPaginatedResponse(queryOptions.Page, queryOptions.PageSize, int(count), expenseCategories)
	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) UpdateExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	log.Println("Updating an expenseCategory")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid expenseCategory ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	expenseCategory := new(models.ExpenseCategory)
	if err := utils.ParseAndValidateData(ctx, expenseCategory); err != nil {
		response := h.responseUtil.SendErrorResponse(err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	expenseCategory.ID = uint(id)
	if err := h.expenseCategoriesService.UpdateCategory(expenseCategory); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to update expenseCategory", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) DeleteExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	log.Println("Deleting an expenseCategory")
	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid expenseCategory ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	_id := uint(id)
	if err := h.expenseCategoriesService.DeleteCategory(_id); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to delete expenseCategory", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}
