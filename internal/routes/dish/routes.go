package dishroutes

import (
	dishhandlers "2kitchen/internal/handlers/dish"

	"github.com/gofiber/fiber/v2"
)

func SetupDishRoutes(app *fiber.App, h *dishhandlers.DishHandler) {
	dishesGroup := app.Group("/dishes")
	dishesGroup.Get("", h.AllDishes)
	dishesGroup.Get("/:restId/:id", h.RestaurantDish)
	dishesGroup.Get("/:restId", h.RestaurantDishes)
	dishesGroup.Post("", h.AddRestaurantDish)
	dishesGroup.Delete("", h.RemoveRestaurantDish)
}
