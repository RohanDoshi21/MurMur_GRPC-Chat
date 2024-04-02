package routes

import "github.com/gofiber/fiber/v2"

func RouteSetup(app *fiber.App) {
	testGroup := app.Group("/test")
	setupTestRoutes(testGroup)

	messageGroup := app.Group("/message")
	setupMessageRoutes(messageGroup)

	userGroup := app.Group("/user")
	setupUserRoutes(userGroup)
}
