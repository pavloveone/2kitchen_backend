package userroutes

import (
	userhandlers "2kitchen/internal/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *userhandlers.UserHander) {
	userGroup := app.Group("/users")
	userGroup.Get("", h.GetAllUsers)
	userGroup.Get("/:id", h.GetUserById)
	userGroup.Post("", h.AddUser)
	userGroup.Post("/login", h.LogIn)
}
