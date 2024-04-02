package main

//go:generate sqlboiler --wipe --no-tests psql

import (
	"os"

	"github.com/RohanDoshi21/messaging-platform/db"
	"github.com/RohanDoshi21/messaging-platform/middleware"
	"github.com/RohanDoshi21/messaging-platform/routes"
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

	casdoorsdk.InitConfig("http://localhost:9080", "a49def21d542e3f7bfb9", "5db49357cc6e7d07427c17b9004f642ecedacc7e", certificate, "message-org", "message-app")

	db.InitDB()

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(requestid.New())
	app.Use(middleware.Auth)
	routes.RouteSetup(app)

	app.Listen(":5000")
}
