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
}

func NewExpenseHandler(service services.ExpenseService) ExpenseHandler {
	return &expenseHandler{
		expenseService: service,
	}
}

func (h *expenseHandler) CreateExpenseHandler(ctx *fiber.Ctx) error {
	println("Handling creation of expense:")
	expense := new(models.CreateExpenseDto)
	if err := ctx.BodyParser(expense); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse expense data",
		})
	}

	if expense.Amount <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Expense amount must be greater than zero",
		})

	}

	if expense.Description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Expense description cannot be empty",
		})
	}

	println("Creating expense with amount:", expense.Amount, "and description:", expense.Description)
	if err := h.expenseService.CreateExpense(expense); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create expense",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(expense)
}

func (h *expenseHandler) GetExpenseByIDHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense ID",
		})
	}

	println("Fetching expense with ID:", id)
	expense, err := h.expenseService.GetExpenseByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expense",
		})
	}

	if expense == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Expense not found",
		})
	}

	return ctx.JSON(expense)
}

func (h *expenseHandler) GetAllExpensesHandler(ctx *fiber.Ctx) error {
	println("Fetching all expenses")
	queryOptions := &models.QueryOptions{
		Query:    ctx.Query("query", ""),
		Page:     ctx.QueryInt("page", 1),
		PageSize: ctx.QueryInt("page_size", 10),
	}

	expenses, count, err := h.expenseService.GetAllExpenses(queryOptions)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expenses",
		})
	}

	apiRepose := utils.NewApiResponse[*models.Expense]()
	response := apiRepose.SendPaginatedResponse(queryOptions.Page, queryOptions.PageSize, int(count), expenses)

	return ctx.JSON(response)
}

func (h *expenseHandler) UpdateExpenseHandler(ctx *fiber.Ctx) error {
	apiRepose := utils.NewApiResponse[any]()
	log.Println("Updating an expense")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := apiRepose.SendErrorResponse("Invalid expense ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	expense := new(models.CreateExpenseDto)
	if err := ctx.BodyParser(expense); err != nil {
		response := apiRepose.SendErrorResponse("Failed to parse expense data", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	log.Printf("Updating expense with ID: %+v", id)
	log.Printf("New Data: %+v", expense)
	if err := h.expenseService.UpdateExpense(id, expense); err != nil {
		response := apiRepose.SendErrorResponse("Failed to update expense", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := apiRepose.SendNoContedResponse()
	return ctx.JSON(response)
}

func (h *expenseHandler) DeleteExpenseHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense ID",
		})
	}

	println("Deleting expense with ID:", id)
	if err := h.expenseService.DeleteExpense(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete expense",
		})
	}

	apiRepose := utils.NewApiResponse[any]()
	response := apiRepose.SendNoContedResponse()

	return ctx.JSON(response)
}
