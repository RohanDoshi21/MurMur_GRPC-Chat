package routes

import (
	"github.com/RohanDoshi21/messaging-platform/api/controller"
	mw "github.com/RohanDoshi21/messaging-platform/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupMessageRoutes(router fiber.Router) {
	// GET my messages
	router.Get("/", mw.Transaction, controller.GetUserMessages)

	// POST a message to a particular user
	router.Post("/", mw.Transaction, controller.SendMessage)

}
