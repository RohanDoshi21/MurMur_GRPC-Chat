package handler

import (
	U "github.com/RohanDoshi21/messaging-platform/util"
	"github.com/gofiber/fiber/v2"
)

// Rollback the transaction from the context
func RollbackCtxTrx(ctx *fiber.Ctx) {
	trx := U.GetPGTrxFromFiberCtx(ctx)
	if trx != nil {
		trx.Rollback()
	}
}

// Commits the transaction from the context
func CommitCtxTrx(ctx *fiber.Ctx) error {
	trx := U.GetPGTrxFromFiberCtx(ctx)
	if trx != nil {
		err := trx.Commit()
		if err != nil {
			return BuildError(ctx, "Error while committing the transaction!", fiber.ErrInternalServerError.Code, err)
		}
	}

	return nil
}
