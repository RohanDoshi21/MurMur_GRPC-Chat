package main

//go:generate sqlboiler --wipe --no-tests psql

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/RohanDoshi21/messaging-platform/api/service"
	S "github.com/RohanDoshi21/messaging-platform/api/service"
	pb "github.com/RohanDoshi21/messaging-platform/proto"
	lgr "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

var (
	doneAPI  = make(chan bool)
	doneGRPC = make(chan bool)
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

	go func() {
		app := fiber.New(fiber.Config{
			Immutable:    true,
			ErrorHandler: handler.ErrorHandler,
		})

		app.Use(cors.New())

		app.Use(requestid.New())

		routes.RouteSetup(app)

		app.Listen(":" + fmt.Sprint(configValues.APP_PORT))
		fmt.Println("Server Starting on Port " + ":" + fmt.Sprint(configValues.APP_PORT))

		doneAPI <- true
	}()

	go func() {
		// Start the server
		// GRPC Server
		server := &S.GrpcServer{
			Clients: make(map[string]pb.GrpcServerService_SendMessageServer),
		}
		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(server.UnaryAuthInterceptor),
			grpc.StreamInterceptor(server.StreamAuthInterceptor),
		)
		pb.RegisterGrpcServerServiceServer(grpcServer, server)
		reflection.Register(grpcServer)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configValues.GRPC_PORT))
		if err != nil {
			log.Fatal("Error creating server", err)
		}
		log.Printf("gRPC server listening on %s", fmt.Sprintf(":%d", configValues.GRPC_PORT))
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("Error serving gRPC server", err)
		}

		doneGRPC <- true
	}()

	service.InitMQTT()

	<-doneAPI
	<-doneGRPC
}
