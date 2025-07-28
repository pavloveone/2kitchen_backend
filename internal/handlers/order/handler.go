package orderhandlers

import (
	"2kitchen/internal/models"
	orderservices "2kitchen/internal/services/order"
	"context"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *orderservices.OrderService
	ctx     context.Context
}

func NewOrderHandler(service *orderservices.OrderService, ctx context.Context) *OrderHandler {
	return &OrderHandler{service: service, ctx: ctx}
}

func (h *OrderHandler) AllOrders(c *fiber.Ctx) error {
	orders, _ := h.service.AllOrders(h.ctx)
	return c.Status(fiber.StatusOK).JSON(orders)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	req := models.CreateOrder{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	id, err := h.service.CreateOrder(h.ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
