package main

import (
	dishhandlers "2kitchen/internal/handlers/dish"
	orderhandlers "2kitchen/internal/handlers/order"
	dishrepositories "2kitchen/internal/repositories/dish"
	orderrepositories "2kitchen/internal/repositories/order"
	dishroutes "2kitchen/internal/routes/dish"
	orderroutes "2kitchen/internal/routes/order"
	dishservices "2kitchen/internal/services/dish"
	orderservices "2kitchen/internal/services/order"

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
	rDishes, err := dishrepositories.NewDishRepository("dishes.db")
	if err != nil {
		logrus.Fatal("Error initializing dishes repository:", err)
	}
	sDishes := dishservices.NewDishService(rDishes)
	hDishes := dishhandlers.NewDishHandler(sDishes)
	dishroutes.SetupDishRoutes(app, hDishes)

	// orders
	rOrders, err := orderrepositories.NewOrderRepository("orders.db")
	if err != nil {
		logrus.Fatal("Error initializing orders repository:", err)
	}
	sOrders := orderservices.NewOrderService(rOrders)
	hOrders := orderhandlers.NewOrderHandler(sOrders)
	orderroutes.SetupOrderRoutes(app, hOrders)

	port := "8080"
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server starting on port")
	logrus.Fatal(app.Listen(":" + port))
}
