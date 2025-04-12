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
	repo, _ := repositories.NewDishRepository("dishes.db")
	service := services.NewDishService(repo)
	handler := handlers.NewDishHandler(service)
	routes.SetupDishRoutes(app, handler)
	port := "80"
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server starting on port")
	logrus.Fatal(app.Listen(":" + port))
}
