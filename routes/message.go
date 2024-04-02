package routes

import (
	"context"
	"log"

	"github.com/RohanDoshi21/messaging-platform/db"
	"github.com/RohanDoshi21/messaging-platform/models"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupMessageRoutes(router fiber.Router) {
	// GET my messages
	router.Get("/", func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(casdoorsdk.User)
		if !ok {
			log.Println("User not found")
		}

		messages, err := models.Messages(models.MessageWhere.Receiver.EQ(user.Id)).All(c.Context(), db.CONN)
		if err != nil {
			// Handle error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch messages",
			})
		}

		// TODO: Implement a way to get the sender's name and email
		// for _, message := range messages {
		// 	sender, err := casdoorsdk.GetUser(message.Sender)
		// 	if err != nil {
		// 		// Handle error
		// 	}
		// 	message.SenderName = sender.Name
		// 	message.SenderEmail = sender.Email
		// }

		// Send the messages as the response.
		return c.JSON(messages)
	})

	// GET messages from a particular user

	// POST a message to a particular user
	router.Post("/", func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(casdoorsdk.User)
		if !ok {
			log.Println("User not found")
		}

		message := new(models.Message)
		if err := c.BodyParser(message); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse the request body",
			})
		}

		message.Sender = user.Id
		err := message.Insert(context.Background(), db.CONN, boil.Infer())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to insert the message into the database",
			})
		}

		// Send the message as the response.
		return c.JSON(message)
	})

}
