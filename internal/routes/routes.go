package routes

import (
	"fiber-test/internal/handlers"
	"fiber-test/internal/middlewares"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middlewares.CheckMiddle)
	app.Get("/docs/*", swaggo.HandlerDefault)

	// Public Routes
	app.Post("/login", handlers.UserLogin)

	api := app.Group("/api/v1", middlewares.Protected())
	//Admin Routes
	api.Get("/names", middlewares.PermissionAdmin, handlers.GetAllNames)
	api.Get("/names/:id", middlewares.PermissionAdmin, handlers.GetNameById)
	api.Post("/names", middlewares.PermissionAdmin, handlers.CreateName)
	api.Put("/names/:id", middlewares.PermissionAdmin, handlers.UpdateName)
	api.Delete("/names/:id", middlewares.PermissionAdmin, handlers.DeleteName)
	api.Post("/upload", middlewares.PermissionAdmin, handlers.UploadFile)
}
