package routes

import (
	"2kitchen/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupDishRoutes(app *fiber.App, dishHandler *handlers.DishHandler) {
	dishesGroup := app.Group("/dishes")
	dishesGroup.Get("", dishHandler.AllDishes)
	dishesGroup.Get("/:restId", dishHandler.RestaurantDishes)
	dishesGroup.Get("/:restId/:id", dishHandler.RestaurantDish)
	dishesGroup.Post("", dishHandler.AddRestaurantDish)
}
