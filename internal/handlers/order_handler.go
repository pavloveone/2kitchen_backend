package handlers

import (
	"2kitchen/internal/models"
	"2kitchen/internal/services"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) AllOrders(c *fiber.Ctx) error {
	orders, _ := h.service.AllOrders()
	return c.Status(fiber.StatusOK).JSON(orders)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	req := models.CreateOrder{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	id, err := h.service.CreateOrder(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}
