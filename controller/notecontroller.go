package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"go-restapi-fiber-lem/data/request"
	"go-restapi-fiber-lem/data/response"
	"go-restapi-fiber-lem/helpers"
	"go-restapi-fiber-lem/service"
	"strconv"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (controller *NoteController) Create(ctx fiber.Ctx) error {
	createNoteRequest := request.CreateNoteRequest{}
	err := json.Unmarshal(ctx.Body(), &createNoteRequest)
	helpers.ErrorPanic(err)

	controller.noteService.Create(createNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(ctx fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	err := json.Unmarshal(ctx.Body(), &updateNoteRequest)
	helpers.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helpers.ErrorPanic(err)
	updateNoteRequest.Id = id

	controller.noteService.Update(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) Delete(ctx fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helpers.ErrorPanic(err)

	controller.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted notes data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindById(ctx fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helpers.ErrorPanic(err)

	noteResponse := controller.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get notes data by Id!",
		Data:    noteResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindAll(ctx fiber.Ctx) error {
	noteResponse := controller.noteService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully get all notes data!",
		Data:    noteResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
