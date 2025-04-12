package handlers

import (
	"2kitchen/internal/models"
	"2kitchen/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DishHandler struct {
	service *services.DishService
}

func NewDishHandler(service *services.DishService) *DishHandler {
	return &DishHandler{service: service}
}

func (h *DishHandler) AllDishes(c *fiber.Ctx) error {
	dishes, _ := h.service.GetAllDishes()
	var code int
	if len(dishes) == 0 {
		code = fiber.StatusNoContent
	} else {
		code = fiber.StatusOK
	}
	return c.Status(code).JSON(dishes)
}

func (h *DishHandler) RestaurantDishes(c *fiber.Ctx) error {
	idParam := c.Params("restId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	dishes, err := h.service.GetRestDishes(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "dishes not found"})
	}
	var code int
	if len(dishes) == 0 {
		code = fiber.StatusNoContent
	} else {
		code = fiber.StatusOK
	}
	return c.Status(code).JSON(dishes)
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
	dish, err := h.service.DishById(restId, id)
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
	dishId, err := h.service.AddDish(req)
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
	err := h.service.RemoveDish(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.SendStatus(fiber.StatusOK)
}
