package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	port := "80"
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Server starting on port")
	logrus.Fatal(app.Listen(":" + port))
}
