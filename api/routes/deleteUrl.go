package routes

import (
	"github.com/ThisIsNotJustin/goshort/api/database"
	"github.com/gofiber/fiber/v2"
)

func DeleteURL(c *fiber.Ctx) error {
	shortID := c.Params("shortID")

	err := database.Client.Del(database.Ctx, shortID).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to delete link",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted shortened URL",
	})
}
