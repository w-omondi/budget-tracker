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
}

func NewExpenseCategoriesHandler(service services.ExpenseCategoryService) ExpenseCategoriesHandler {
	return &expenseCategoriesHandler{
		expenseCategoriesService: service,
	}
}

func (h *expenseCategoriesHandler) CreateExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	println("Handling creation of expenseCategory:")
	expenseCategory := new(models.CreateExpenseCategoryDto)
	if err := ctx.BodyParser(expenseCategory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse expenseCategory data",
		})
	}

	if expenseCategory.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Expense category name cannot be empty",
		})
	}

	if err := h.expenseCategoriesService.CreateCategory(expenseCategory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create expenseCategory",
		})
	}

	apiRepose := utils.NewApiResponse[any]()
	response := apiRepose.SendNoContedResponse()

	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) GetExpenseCategoryByIDHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expenseCategory ID",
		})
	}

	println("Fetching expenseCategory with ID:", id)
	expenseCategory, err := h.expenseCategoriesService.GetCategoryByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expenseCategory",
		})
	}

	if expenseCategory == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Expense Category not found",
		})
	}

	res := utils.NewApiResponse[models.ExpenseCategory]()
	response := res.SendSingleResponse(expenseCategory)

	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) GetAllExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	println("Fetching all expenseCategories")
	queryOptions := &models.QueryOptions{
		Query:    ctx.Query("query", ""),
		Page:     ctx.QueryInt("page", 1),
		PageSize: ctx.QueryInt("page_size", 10),
	}

	expenseCategories, count, err := h.expenseCategoriesService.GetAllCategories(queryOptions)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expenseCategories",
		})
	}

	apiRepose := utils.NewApiResponse[*models.ExpenseCategory]()
	response := apiRepose.SendPaginatedResponse(queryOptions.Page, queryOptions.PageSize, int(count), expenseCategories)

	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) UpdateExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expenseCategory ID",
		})
	}

	expenseCategory := new(models.ExpenseCategory)
	if err := ctx.BodyParser(expenseCategory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse expenseCategory data",
		})
	}

	expenseCategory.ID = uint(id)
	log.Println("Updating expenseCategory with ID:", id)
	log.Printf("New Data: %+v", expenseCategory)
	if err := h.expenseCategoriesService.UpdateCategory(expenseCategory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update expenseCategory",
		})
	}

	apiRepose := utils.NewApiResponse[any]()
	response := apiRepose.SendNoContedResponse()

	return ctx.JSON(response)
}

func (h *expenseCategoriesHandler) DeleteExpenseCategoriesHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expenseCategory ID",
		})
	}

	_id := uint(id)

	println("Deleting expenseCategory with ID:", _id)
	if err := h.expenseCategoriesService.DeleteCategory(_id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete expenseCategory",
		})
	}

	apiRepose := utils.NewApiResponse[any]()
	response := apiRepose.SendNoContedResponse()

	return ctx.JSON(response)
}
