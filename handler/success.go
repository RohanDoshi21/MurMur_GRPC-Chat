package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Sends a success (200) response for the data to the user.
func Success(ctx *fiber.Ctx, data interface{}) error {
	err := CommitCtxTrx(ctx)

	if err != nil {
		return err
	}

	return ctx.JSON(data)
}
