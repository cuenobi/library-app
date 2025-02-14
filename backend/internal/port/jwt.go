package port

import "github.com/gofiber/fiber/v2"

type JWT interface {
	Generate(username, role string) string
	Validate(c *fiber.Ctx) error
}
