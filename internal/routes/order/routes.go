package orderroutes

import (
	_ "2kitchen/internal/auth"
	orderhandler "2kitchen/internal/handlers/order"

	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(app *fiber.App, h *orderhandler.OrderHandler) {
	ordersGroup := app.Group("/orders")
	ordersGroup.Get("", h.AllOrders)
	ordersGroup.Post("", h.CreateOrder)
}
