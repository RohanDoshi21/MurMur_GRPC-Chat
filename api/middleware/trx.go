package middleware

import (
	"context"

	DB "github.com/RohanDoshi21/messaging-platform/db"
	H "github.com/RohanDoshi21/messaging-platform/handler"
	"github.com/gofiber/fiber/v2"
)

// Injects a new transaction to the context
func Transaction(ctx *fiber.Ctx) error {
	dbCtx := context.Background()
	pgTrx, err := DB.PGTransaction(dbCtx)

	if err != nil {
		return H.BuildError(ctx, "Failed to initiate a transaction!", fiber.ErrInternalServerError.Code, err)
	}

	ctx.Locals("pgTrx", pgTrx)

	ctx.Next()

	return nil
}
