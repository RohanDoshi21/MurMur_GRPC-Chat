package routes

import (
	"context"

	"github.com/RohanDoshi21/messaging-platform/db"
	"github.com/RohanDoshi21/messaging-platform/models"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupTestRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.Get("/db", func(c *fiber.Ctx) error {
		message := &models.Message{
			Sender:   "Rohan",
			Receiver: "Rohit",
			Content:  "Hello, World!",
		}
		err := message.Insert(context.Background(), db.CONN, boil.Infer())
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendString("Connected to DB!")
	})
}
