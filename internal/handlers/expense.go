package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
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
	expenses, err := h.expenseService.GetAllExpenses()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expenses",
		})
	}

	return ctx.JSON(expenses)
}

func (h *expenseHandler) UpdateExpenseHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense ID",
		})
	}

	expense := new(models.Expense)
	if err := ctx.BodyParser(expense); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse expense data",
		})
	}

	expense.ID = id
	println("Updating expense with ID:", id)
	if err := h.expenseService.UpdateExpense(expense); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update expense",
		})
	}

	return ctx.JSON(expense)
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

	return ctx.SendStatus(fiber.StatusNoContent)
}
