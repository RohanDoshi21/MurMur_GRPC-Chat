package middleware

import (
	H "github.com/RohanDoshi21/messaging-platform/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(body any) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		if err := ctx.BodyParser(body); err != nil {
			msg := "Failed to parse the body!"
			return H.BuildError(ctx, msg, fiber.ErrBadRequest.Code, err)
		}
		validate := validator.New()
		err := validate.Struct(body)
		if err != nil {
			return H.BuildError(ctx, err.Error(), fiber.ErrBadRequest.Code, err)
		} else {
			ctx.Locals("body", body)
		}
		ctx.Next()
		return nil

	}

}
