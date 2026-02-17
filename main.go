package main

import (
	"log"

	_ "fiber-test/docs"
	"fiber-test/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

// @title           Fiber Test API
// @version         1.0
// @description     ระบบทดสอบ API ด้วย Fiber v3
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description พิมพ์คำว่า "Bearer " ตามด้วย Token (เช่น "Bearer eyJhb...")
func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET,POST,HEAD,PUT,DELETE,PATCH"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
	}))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
