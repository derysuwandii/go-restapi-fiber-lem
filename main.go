package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"go-restapi-fiber-lem/configs"
	"go-restapi-fiber-lem/controller"
	"go-restapi-fiber-lem/models"
	"go-restapi-fiber-lem/repository"
	"go-restapi-fiber-lem/router"
	"go-restapi-fiber-lem/service"
	"log"
)

func main() {
	log.Println("Run service on port 8080")

	loadConfig, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env ", err)
	}

	db := configs.ConnectDB(loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&models.Note{})

	//Init repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	//Init service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	//init controller
	noteController := controller.NewNoteController(noteService)

	//Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Use("/api", routes)

	log.Fatal(app.Listen(":8080"))
}
