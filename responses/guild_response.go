package responses

import "github.com/gofiber/fiber/v2"

type GuildResponse struct {
	// TODO: Finish and complete the response type for when users grab/post to the /guild API route

	Data *fiber.Map `json:"data"`
}