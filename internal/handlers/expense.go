package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
)

var expenseService *services.ExpenseService

func CreateExpenseHandler(ctx *fiber.Ctx) error {
	println("Handling creation of expense:")
	expense := new(models.Expense)
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
	if err := expenseService.CreateExpense(expense); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create expense",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(expense)
}

func GetExpenseByIDHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense ID",
		})
	}

	println("Fetching expense with ID:", id)
	expense, err := expenseService.GetExpenseByID(id)
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

func GetAllExpensesHandler(ctx *fiber.Ctx) error {
	println("Fetching all expenses")
	expenses, err := expenseService.GetAllExpenses()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expenses",
		})
	}

	return ctx.JSON(expenses)
}

func UpdateExpenseHandler(ctx *fiber.Ctx) error {
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
	if err := expenseService.UpdateExpense(expense); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update expense",
		})
	}

	return ctx.JSON(expense)
}

func DeleteExpenseHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense ID",
		})
	}

	println("Deleting expense with ID:", id)
	if err := expenseService.DeleteExpense(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete expense",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
