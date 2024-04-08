package util

import (
	"database/sql"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Gets the authenticated user model
func GetAuthUser(ctx *fiber.Ctx) casdoorsdk.User {
	return ctx.Locals("user").(casdoorsdk.User)
}

func GetPGTrxFromFiberCtx(ctx *fiber.Ctx) *sql.Tx {
	trxInf := ctx.Locals("pgTrx")

	if trxInf == nil {
		return nil
	}

	return ctx.Locals("pgTrx").(*sql.Tx)
}
func UUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
