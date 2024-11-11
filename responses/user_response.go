package responses

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	// TODO: Finish the response for when developers grab information or try posting information to the API

	Data *fiber.Map `json:"data"`
}