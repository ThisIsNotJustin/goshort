package routes

import (
	"github.com/ThisIsNotJustin/goshort/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetByShortID(c *fiber.Ctx) error {
	shortID := c.Params("shortID")

	val, err := database.Client.Get(database.Ctx, shortID).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data not found for given ShortID",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect to DB",
		})
	}

	return c.Redirect(val, fiber.StatusMovedPermanently)
}
