package main

import (
	"2kitchen/internal/handlers"
	"2kitchen/internal/repositories"
	"2kitchen/internal/routes"
	"2kitchen/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://2kitchen-frontend.vercel.app",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// dishes
	rDishes, err := repositories.NewDishRepository("dishes.db")
	if err != nil {
		logrus.Fatal("Error initializing dishes repository:", err)
	}
	sDishes := services.NewDishService(rDishes)
	hDishes := handlers.NewDishHandler(sDishes)
	routes.SetupDishRoutes(app, hDishes)

	// orders
	rOrders, err := repositories.NewOrderRepository("orders.db")
	if err != nil {
		logrus.Fatal("Error initializing orders repository:", err)
	}
	sOrders := services.NewOrderService(rOrders)
	hOrders := handlers.NewOrderHandler(sOrders)
	routes.SetupOrderRoutes(app, hOrders)

	port := "8080"
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server starting on port")
	logrus.Fatal(app.Listen(":" + port))
}
