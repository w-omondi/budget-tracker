package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"github.com/w-omondi/budget-tracker.git/internal/utils"
)

type DashboardHandler interface {
	GetDashbordSummary(ctx *fiber.Ctx) error
}

type DashBoordResponseData struct {
	Expenses   []*models.Expense         `json:"expenses"`
	Revenues   []*models.Revenue         `json:"revenues"`
	Categories []*models.ExpenseCategory `json:"categories"`
}

type dashboardHandler struct {
	expenseService  services.ExpenseService
	revenueService  services.RevenueService
	categoryService services.ExpenseCategoryService
	responseUtil    utils.ApiResponse[*DashBoordResponseData]
}

func NewDashbordHandler(
	expenseService services.ExpenseService,
	revenueService services.RevenueService,
	categoryService services.ExpenseCategoryService,
) DashboardHandler {
	return &dashboardHandler{
		expenseService:  expenseService,
		revenueService:  revenueService,
		categoryService: categoryService,
		responseUtil:    utils.NewApiResponse[*DashBoordResponseData](),
	}
}

func (h *dashboardHandler) GetDashbordSummary(ctx *fiber.Ctx) error {
	log.Println("Fetching dashboard data")

	queryOptions := &models.QueryOptions{
		Page:     1,
		PageSize: 20,
		Query:    "",
	}

	expenses, _, err := h.expenseService.GetAllExpenses(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Unable to retrive expenses", err)
		return ctx.JSON(response)
	}

	revenues, _, err := h.revenueService.GetAllRevenues(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Unable to retrive revenues", err)
		return ctx.JSON(response)
	}

	categories, _, err := h.categoryService.GetAllCategories(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Unable to retrive revenues", err)
		return ctx.JSON(response)
	}

	aggregatedResponse := &DashBoordResponseData{
		Expenses:   expenses,
		Revenues:   revenues,
		Categories: categories,
	}

	response := h.responseUtil.SendSingleResponse(aggregatedResponse)
	return ctx.JSON(response)
}
