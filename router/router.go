package router

import (
	"github.com/gofiber/fiber/v3"
	"go-restapi-fiber-lem/controller"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcode Golang",
		})
	})

	router.Post("/notes", noteController.Create)
	router.Get("/notes", noteController.FindAll)
	router.Delete("/notes/:noteId", noteController.Delete)
	router.Get("/notes/:noteId", noteController.FindById)
	router.Patch("/notes/:noteId", noteController.Update)

	return router
}
