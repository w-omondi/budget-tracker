package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"github.com/w-omondi/budget-tracker.git/internal/utils"
)

type RevenueHandler interface {
	CreateRevenueHandler(ctx *fiber.Ctx) error
	GetRevenueByIDHandler(ctx *fiber.Ctx) error
	GetAllRevenuesHandler(ctx *fiber.Ctx) error
	UpdateRevenueHandler(ctx *fiber.Ctx) error
	DeleteRevenueHandler(ctx *fiber.Ctx) error
}

type revenueHandler struct {
	revenueService services.RevenueService
	responseUtil   utils.ApiResponse[*models.Revenue]
}

func NewRevenueHandler(service services.RevenueService) RevenueHandler {
	return &revenueHandler{
		revenueService: service,
		responseUtil:   utils.NewApiResponse[*models.Revenue](),
	}
}

func (h *revenueHandler) CreateRevenueHandler(ctx *fiber.Ctx) error {
	log.Println("Handling creation of revenue:")

	revenue := new(models.CreateRevenueDto)
	if err := utils.ParseAndValidateData(ctx, revenue); err != nil {
		response := h.responseUtil.SendErrorResponse(err.Error(), err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.revenueService.CreateRevenue(revenue); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to create revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	log.Println("Revenue created successfully")
	return ctx.Status(fiber.StatusCreated).JSON(revenue)
}

func (h *revenueHandler) GetRevenueByIDHandler(ctx *fiber.Ctx) error {
	log.Println("Fetching revenue by ID")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid revenue ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	revenue, err := h.revenueService.GetRevenueByID(uint(id))
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to fetch revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if revenue == nil {
		response := h.responseUtil.SendErrorResponse("Revenue not found", nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}

	response := h.responseUtil.SendSingleResponse(revenue)
	return ctx.JSON(response)
}

func (h *revenueHandler) GetAllRevenuesHandler(ctx *fiber.Ctx) error {
	log.Println("Fetching all revenues")

	queryOptions := &models.QueryOptions{
		Query:    ctx.Query("query", ""),
		Page:     ctx.QueryInt("page", 1),
		PageSize: ctx.QueryInt("page_size", 10),
	}

	revenues, count, err := h.revenueService.GetAllRevenues(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to fetch revenues", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendPaginatedResponse(queryOptions.Page, queryOptions.PageSize, int(count), revenues)
	return ctx.JSON(response)
}

func (h *revenueHandler) UpdateRevenueHandler(ctx *fiber.Ctx) error {
	log.Println("Updating an revenue")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid revenue ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	revenue := new(models.CreateRevenueDto)
	if err := utils.ParseAndValidateData(ctx, revenue); err != nil {
		response := h.responseUtil.SendErrorResponse(err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.revenueService.UpdateRevenue(uint(id), revenue); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to update revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}

func (h *revenueHandler) DeleteRevenueHandler(ctx *fiber.Ctx) error {
	log.Println("Deleting an revenue")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid revenue ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.revenueService.DeleteRevenue(uint(id)); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to delete revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.JSON(response)
}

