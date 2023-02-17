package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals("DB", db)
		return c.Next()
	}
}
