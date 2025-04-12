package routes

import (
	"2kitchen/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(app *fiber.App, h *handlers.OrderHandler) {
	ordersGroup := app.Group("/orders")
	ordersGroup.Get("", h.AllOrders)
	ordersGroup.Post("", h.CreateOrder)
}
