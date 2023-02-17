package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, userService UserService) {
	// Ruta para crear un nuevo usuario
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return err
		}

		if err := userService.Create(user); err != nil {
			return err
		}

		return c.JSON(user)
	})

	// Ruta para obtener un usuario por ID
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user, err := userService.FindByID(id)
		if err != nil {
			return err
		}

		return c.JSON(user)
	})

	// Ruta para obtener todos los usuarios
	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := userService.FindAll()
		if err != nil {
			return err
		}

		return c.JSON(users)
	})
}
