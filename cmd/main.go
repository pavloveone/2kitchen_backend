package main

import (
	"2kitchen/internal/handlers"
	"2kitchen/internal/repositories"
	"2kitchen/internal/routes"
	"2kitchen/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	// dishes
	rDishes, _ := repositories.NewDishRepository("dishes.db")
	sDishes := services.NewDishService(rDishes)
	hDishes := handlers.NewDishHandler(sDishes)
	routes.SetupDishRoutes(app, hDishes)

	// orders
	rOrders, _ := repositories.NewOrderRepository("orders.db")
	sOrders := services.NewOrderService(rOrders)
	hOrders := handlers.NewOrderHandler(sOrders)
	routes.SetupOrderRoutes(app, hOrders)

	port := "80"
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server starting on port")
	logrus.Fatal(app.Listen(":" + port))
}
