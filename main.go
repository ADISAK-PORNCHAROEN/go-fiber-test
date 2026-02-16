package main

import (
	namesService "fiber-test/internal"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

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

	secretKey := os.Getenv("SECRET_KEY")
	println(secretKey)

	namesService.NameList = append(namesService.NameList, namesService.Name{Id: 1, Fname: "Adisak", Lname: "Porncharoen"})
	namesService.NameList = append(namesService.NameList, namesService.Name{Id: 2, Fname: "Mathawee", Lname: "Pumpuang"})

	app.Get("/", namesService.HelloWorld)
	app.Get("/Names", namesService.Names)
	app.Get("/Names/:id", namesService.NamesId)
	app.Post("/Names", namesService.CreateName)
	app.Put("/Names/:id", namesService.UpdateName)
	app.Delete("/Names/:id", namesService.DeleteName)

	app.Post("/upload", namesService.UploadFile)

	app.Listen(":8080")
}
