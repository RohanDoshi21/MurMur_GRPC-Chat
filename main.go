package main

//go:generate sqlboiler --wipe --no-tests psql

import (
	"fmt"
	"os"

	lgr "github.com/sirupsen/logrus"

	"github.com/RohanDoshi21/messaging-platform/api/routes"
	C "github.com/RohanDoshi21/messaging-platform/config"
	"github.com/RohanDoshi21/messaging-platform/db"
	"github.com/RohanDoshi21/messaging-platform/handler"
	cSdk "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

// Gives the API server certificate bytes from the local filesystem (TODO: DB impl)
func GetAPIServerCert(config C.Config) ([]byte, error) {
	serverCertLoc := config.CASDOOR_CERTIFICATE
	cert, err := os.ReadFile(serverCertLoc)

	if err != nil {
		return nil, err
	}

	return cert, nil
}

// Initiates casdoor global authentication configuration.
func InitAuthConfig(config C.Config) error {
	serverCertBytes, err := GetAPIServerCert(config)

	if err != nil {
		return err
	}

	casdoorEndpoint := config.CASDOOR_ENDPOINT
	clientId := config.CASDOOR_CLIENT_ID
	clientSecret := config.CASDOOR_CLIENT_SECRET
	casdoorOrganization := config.CASDOOR_ORG_NAME
	casdoorApplication := config.CASDOOR_APP_NAME

	cSdk.InitConfig(casdoorEndpoint, clientId, clientSecret, string(serverCertBytes), casdoorOrganization, casdoorApplication)

	return nil
}

func main() {

	godotenv.Load(".env")

	if _, ok := os.LookupEnv("GO_ENV"); !ok {
		lgr.Fatalln("Error while loading environment variables!")
	}

	configValues, configErr := C.Init()

	if configErr != nil {
		lgr.Fatalln(configErr)
	}

	err := db.Init()

	if err != nil {
		lgr.Fatalln(err)
	}

	defer db.Close()

	err = InitAuthConfig(*configValues)

	if err != nil {
		lgr.Fatalln("Error while setting up auth config!", err)
	}

	app := fiber.New(fiber.Config{
		Immutable:    true,
		ErrorHandler: handler.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length, Authorization",
	}))

	app.Use(requestid.New())

	routes.RouteSetup(app)

	app.Listen(":" + fmt.Sprint(configValues.APP_PORT))
}
