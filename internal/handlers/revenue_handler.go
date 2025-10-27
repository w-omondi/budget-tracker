package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/models"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"github.com/w-omondi/budget-tracker.git/internal/utils"
)

type RevenueHandler interface {
	CreateRevenue(ctx *fiber.Ctx) error
	GetRevenueByID(ctx *fiber.Ctx) error
	GetAllRevenues(ctx *fiber.Ctx) error
	UpdateRevenue(ctx *fiber.Ctx) error
	DeleteRevenue(ctx *fiber.Ctx) error
}

type revenueHandler struct {
	revenueService services.RevenueService
	responseUtil   utils.ApiResponse[*models.Revenue]
}

func NewRevenueHandler(service services.RevenueService) RevenueHandler {
	responseUtil := utils.NewApiResponse[*models.Revenue]()
	return &revenueHandler{
		revenueService: service,
		responseUtil:  responseUtil,
	}
}

func (h *revenueHandler) CreateRevenue(ctx *fiber.Ctx) error {
	revenueDto := new(models.CreateRevenueDto)
	if err := utils.ParseAndValidateData(ctx, revenueDto); err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid input data", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.revenueService.CreateRevenue(revenueDto); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to create revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response:=h.responseUtil.SendSingleResponse(revenueDto)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *revenueHandler) GetRevenueByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid ID parameter", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	revenue, err := h.revenueService.GetRevenueByID(id)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Revenue not found", err)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}

	response := h.responseUtil.SendSingleResponse(revenue)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *revenueHandler) GetAllRevenues(ctx *fiber.Ctx) error {
	queryOptions := utils.ParseQueryOptions(ctx)

	revenues, total, err := h.revenueService.GetAllRevenues(queryOptions)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to retrieve revenues", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendPaginatedResponse(queryOptions.Page, queryOptions.PageSize, int(total), revenues)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *revenueHandler) UpdateRevenue(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid ID parameter", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	revenueDto := new(models.UpdateRevenueDto)
	if err := utils.ParseAndValidateData(ctx, revenueDto); err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid input data", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.revenueService.UpdateRevenue(id, revenueDto); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to update revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendSingleResponse(revenueDto)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *revenueHandler) DeleteRevenue(ctx *fiber.Ctx) error {
	id, err := utils.ParseIDParam(ctx)
	if err != nil {
		response := h.responseUtil.SendErrorResponse("Invalid ID parameter", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.revenueService.DeleteRevenue(uint(id)); err != nil {
		response := h.responseUtil.SendErrorResponse("Failed to delete revenue", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := h.responseUtil.SendNoContedResponse()
	return ctx.Status(fiber.StatusNoContent).JSON(response)
}
	