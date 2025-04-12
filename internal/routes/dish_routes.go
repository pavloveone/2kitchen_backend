package routes

import (
	"2kitchen/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupDishRoutes(app *fiber.App, h *handlers.DishHandler) {
	dishesGroup := app.Group("/dishes")
	dishesGroup.Get("", h.AllDishes)
	dishesGroup.Get("/:restId", h.RestaurantDishes)
	dishesGroup.Get("/:restId/:id", h.RestaurantDish)
	dishesGroup.Post("", h.AddRestaurantDish)
	dishesGroup.Delete("", h.RemoveRestaurantDish)
}
