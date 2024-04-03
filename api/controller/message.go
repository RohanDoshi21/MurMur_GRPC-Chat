package controller

import (
	"context"
	"log"

	S "github.com/RohanDoshi21/messaging-platform/api/service"
	db "github.com/RohanDoshi21/messaging-platform/db"
	H "github.com/RohanDoshi21/messaging-platform/handler"
	"github.com/RohanDoshi21/messaging-platform/models"
	U "github.com/RohanDoshi21/messaging-platform/util"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetUserMessages(c *fiber.Ctx) error {
	user := U.GetAuthUser(c)

	pgTrx := U.GetPGTrxFromFiberCtx(c)

	userMessagesBody := &S.Messages{
		UserID: user.Id,
	}

	messages, err := S.GetUserMessages(userMessagesBody, pgTrx)

	if err != nil {
		// Handle error
		return H.BuildError(c, err.Message, err.Code, err.Error)
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
	return H.Success(c, fiber.Map{
		"ok":       1,
		"messages": messages,
	})
}

func SendMessage(c *fiber.Ctx) error {
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
	err := message.Insert(context.Background(), db.PostgresConn, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert the message into the database",
		})
	}

	// Send the message as the response.
	return c.JSON(message)
}
