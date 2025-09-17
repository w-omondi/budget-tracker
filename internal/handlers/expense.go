package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"github.com/w-omondi/budget-tracker.git/internal/utils"
)

type ExpenseHandler interface {
	CreateExpenseHandler(ctx *fiber.Ctx) error
	GetExpenseByIDHandler(ctx *fiber.Ctx) error
	GetAllExpensesHandler(ctx *fiber.Ctx) error
	UpdateExpenseHandler(ctx *fiber.Ctx) error
	DeleteExpenseHandler(ctx *fiber.Ctx) error
}

type expenseHandler struct {
	expenseService services.ExpenseService
	responseUtil   utils.ApiResponse[*models.Expense]
}

func NewExpenseHandler(service services.ExpenseService) ExpenseHandler {
	return &expenseHandler{
		expenseService: service,
		responseUtil:   utils.NewApiResponse[*models.Expense](),
	}
}

func (h *expenseHandler) CreateExpenseHandler(ctx *fiber.Ctx) error {
	log.Println("Handling creation of expense:")

	expense := new(models.CreateExpenseDto)
	if err := utils.ParseAndValidateData(ctx, expense); err != nil {
		response := h.responseUtil.SendErrorResponse(err.Error(), err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	log.Printf("Creating expense with data: %+v", expense)
	if err := h.expenseService.CreateExpense(expense); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to create expense", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	log.Println("Expense created successfully")
	return ctx.Status(fiber.StatusCreated).JSON(expense)
}

func (h *expenseHandler) GetExpenseByIDHandler(ctx *fiber.Ctx) error {
	log.Println("Fetching expense by ID")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid expense ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	expense, err := h.expenseService.GetExpenseByID(id)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to fetch expense", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if expense == nil {
		response := h.responseUtil.SendErrorResponse("Expense not found", nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}

	return ctx.JSON(expense)
}

func (h *expenseHandler) GetAllExpensesHandler(ctx *fiber.Ctx) error {
	log.Println("Fetching all expenses")

	queryOptions := &models.QueryOptions{
		Query:    ctx.Query("query", ""),
		Page:     ctx.QueryInt("page", 1),
		PageSize: ctx.QueryInt("page_size", 10),
	}

	expenses, count, err := h.expenseService.GetAllExpenses(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to fetch expenses", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendPaginatedResponse(queryOptions.Page, queryOptions.PageSize, int(count), expenses)
	return ctx.JSON(response)
}

func (h *expenseHandler) UpdateExpenseHandler(ctx *fiber.Ctx) error {
	log.Println("Updating an expense")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid expense ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	expense := new(models.CreateExpenseDto)
	if err := utils.ParseAndValidateData(ctx, expense); err != nil {
		response := h.responseUtil.SendErrorResponse(err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.expenseService.UpdateExpense(id, expense); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to update expense", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}

func (h *expenseHandler) DeleteExpenseHandler(ctx *fiber.Ctx) error {
	log.Println("Deleting an expense")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid expense ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.expenseService.DeleteExpense(id); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to delete expense", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}
