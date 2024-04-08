package routes

import (
	"github.com/RohanDoshi21/messaging-platform/api/controller"
	mw "github.com/RohanDoshi21/messaging-platform/api/middleware"
	S "github.com/RohanDoshi21/messaging-platform/api/service"
	"github.com/gofiber/fiber/v2"
)

func setupGroupRoutes(router fiber.Router) {

	router.Get("/:id", mw.Transaction, controller.GetGroupDetails)

	router.Post("/", mw.Transaction, mw.Validate(&S.GroupCreateBody{}), controller.CreateGroup)
	router.Get("/join/:id", mw.Transaction, controller.JoinGroup)

	router.Post("/add", mw.Transaction, mw.Validate(&S.GroupAddUserBody{}), controller.AddUserToGroup)
}
