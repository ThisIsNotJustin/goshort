package routes

import (
	"time"

	"github.com/ThisIsNotJustin/goshort/api/database"
	"github.com/ThisIsNotJustin/goshort/api/models"
	"github.com/gofiber/fiber/v2"
)

func EditURL(c *fiber.Ctx) error {
	shortID := c.Params("shortID")
	var body models.Request

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	val, err := database.Client.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "ShortID does not exist",
		})
	}

	err = database.Client.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update shortened link",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Content Updated",
	})
}
