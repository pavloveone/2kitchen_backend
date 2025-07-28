package dishhandlers

import (
	"2kitchen/internal/models"
	dishservices "2kitchen/internal/services/dish"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DishHandler struct {
	service *dishservices.DishService
	ctx     context.Context
}

func NewDishHandler(service *dishservices.DishService, ctx context.Context) *DishHandler {
	return &DishHandler{service: service, ctx: ctx}
}

func (h *DishHandler) AllDishes(c *fiber.Ctx) error {
	dishes, _ := h.service.GetAllDishes(h.ctx)
	return c.Status(fiber.StatusOK).JSON(dishes)
}

func (h *DishHandler) RestaurantDishes(c *fiber.Ctx) error {
	idParam := c.Params("restId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	dishes, err := h.service.GetRestDishes(h.ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "dishes not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dishes)
}

func (h *DishHandler) RestaurantDish(c *fiber.Ctx) error {
	restParam := c.Params("restId")
	restId, err := strconv.Atoi(restParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid rest ID"})
	}
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid dish ID"})
	}
	dish, err := h.service.DishById(h.ctx, restId, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "dish not found"})
	}
	return c.Status(fiber.StatusOK).JSON(dish)
}

func (h *DishHandler) AddRestaurantDish(c *fiber.Ctx) error {
	req := models.ModificationDish{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	dishId, err := h.service.AddDish(h.ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": dishId})
}

func (h *DishHandler) RemoveRestaurantDish(c *fiber.Ctx) error {
	req := models.ModificationDish{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	err := h.service.RemoveDish(h.ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.SendStatus(fiber.StatusOK)
}
