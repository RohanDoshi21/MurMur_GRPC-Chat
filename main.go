package main

import (
	"os"

	"github.com/RohanDoshi21/messaging-platform/middleware"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	file, err := os.ReadFile("./certs/key.pem")
	if err != nil {
		panic(err)
	}
	certificate := string(file)

	casdoorsdk.InitConfig("http://localhost:9080", "929bbf4c4d484e2e1983", "d523097d12a5bfc6e290a0da77cc89e0718c6101", certificate, "message-app", "message-app-app")

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(requestid.New())
	app.Use(middleware.Auth)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":5000")
}
