package routes

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func setupUserRoutes(router fiber.Router) {

	// GET all users
	router.Get("/", func(c *fiber.Ctx) error {
		users, err := casdoorsdk.GetUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userResponses := make([]UserResponse, len(users))

		for i, user := range users {
			userResponses[i] = UserResponse{
				ID:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			}
		}

		return c.JSON(userResponses)
	})

	// GET a particular user
}
