package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware that receives the error and handles accordingly
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	RollbackCtxTrx(ctx) // Rollback if the context was transactional

	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := "Internal server error!"

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	return ctx.Status(code).JSON(fiber.Map{
		"ok":      0,
		"message": message,
	})
}

// Sends an error message to the client along with logging it!
func BuildError(ctx *fiber.Ctx, message string, code int, originalErr error) error {
	RollbackCtxTrx(ctx) // Rollback if the context was transactional

	if code == 0 {
		code = fiber.ErrBadRequest.Code
	}

	detail := ""

	if originalErr != nil {
		detail = originalErr.Error()
	}

	// TODO: Add remove logging support
	log.Println(message, detail, code)

	return ctx.Status(code).JSON(fiber.Map{
		"ok":      0,
		"message": message,
		"detail":  detail,
	})
}
