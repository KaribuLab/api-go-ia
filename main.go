package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Abrir una conexión a la base de datos SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrar el modelo a la base de datos
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}

	// Crear una nueva instancia de Fiber
	app := fiber.New()

	// Agregar el middleware para la instancia de gorm.DB
	app.Use(DBMiddleware(db))

	// Agregar las rutas a la aplicación
	RegisterRoutes(app, db, NewUserService(NewUserRepository(db)))

	// Iniciar el servidor de Fiber
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
