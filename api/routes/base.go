package routes

import (
	"github.com/RohanDoshi21/messaging-platform/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App) {
	testGroup := app.Group("/test")
	setupTestRoutes(testGroup)

	app.Use(middleware.Auth)

	messageGroup := app.Group("/message")
	setupMessageRoutes(messageGroup)

	userGroup := app.Group("/user")
	setupUserRoutes(userGroup)
}
