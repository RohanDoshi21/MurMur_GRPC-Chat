package middleware

import (
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authHeaders := headers["Authorization"]

	if authHeaders == nil {
		msg := "Authorization header is missing!"
		return &fiber.Error{Code: fiber.ErrBadRequest.Code, Message: msg}
	}

	authHeader := authHeaders[0]

	if authHeader == "" {
		msg := "Authorization header is missing!"
		return &fiber.Error{Code: fiber.ErrBadRequest.Code, Message: msg}
	}

	splitToken := strings.Split(authHeader, " ")

	if len(splitToken) != 2 {
		msg := "Token is missing in the header!"
		return &fiber.Error{Code: fiber.ErrBadRequest.Code, Message: msg}
	}

	authKind := splitToken[0]

	if authKind != "Bearer" {
		msg := "Invalid authorization scheme!"
		return &fiber.Error{Code: fiber.ErrBadRequest.Code, Message: msg}
	}

	token := splitToken[1]
	claims, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		msg := "Invalid token!"
		return &fiber.Error{Code: fiber.ErrBadRequest.Code, Message: msg}
	}

	ctx.Locals("user", claims.User)
	ctx.Locals("user-id", claims.Subject)

	ctx.Next()

	return nil
}
