package controller

import (
	"fmt"
	"strconv"
	"todo-list/model"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type todoItemController struct {
	service service.TodoItemService
}

func NewControllerTodo(service service.TodoItemService) TodoItemController {
	return &todoItemController{
		service: service,
	}
}

func (controller *todoItemController) Route(group fiber.Router) {
	group.Post("", controller.CreateTodoItem)
	group.Get("/:id", controller.GetTodoItem)
	group.Get("", controller.GetAllTodoController)
	group.Get("/:activity_group_id", controller.GetActivityGroupIdController)
	group.Patch("/:id", controller.UpdateTodoItem)
	group.Delete("/:id", controller.DeleteTodoController)
}

func (controller *todoItemController) CreateTodoItem(ctx *fiber.Ctx) error {
	request := new(model.TodoRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(model.Response{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
	}

	response, err := controller.service.Create(*request)
	if err != nil {
		return ctx.Status(500).JSON(model.Response{
			Status:  "Server Error",
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Succsess",
		Message: "Succsess",
		Data:    response,
	})
}

func (controller *todoItemController) GetTodoItem(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	itemID, err := strconv.Atoi(idParam)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", itemID),
			})
		}
	}

	response, err := controller.service.GetOneTodoService(itemID)
	if err != nil {
		return ctx.Status(500).JSON(model.Response{
			Status:  "Server Error",
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Succsess",
		Message: "Succsess",
		Data:    response,
	})
}

func (controller *todoItemController) GetActivityGroupIdController(ctx *fiber.Ctx) error {
	activityGroupIDParam := ctx.Query("activity_group_id")
	activityGroupID, err := strconv.Atoi(activityGroupIDParam)
	if err != nil {
		return ctx.Status(400).JSON(model.Response{
			Status:  "Bad Request",
			Message: "Bad Request",
		})
	}

	response, err := controller.service.GetByActivityGroupID(activityGroupID)
	if err != nil {
		return ctx.Status(500).JSON(model.Response{
			Status:  "Server Error",
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Succsess",
		Message: "Succsess",
		Data:    response,
	})
}

func (controller *todoItemController) GetAllTodoController(ctx *fiber.Ctx) error {
	activityGroupIDParam := ctx.Query("activity_group_id")
	if activityGroupIDParam != "" {
		activityGroupID, err := strconv.Atoi(activityGroupIDParam)
		if err != nil {
			return ctx.Status(400).JSON(model.Response{
				Status:  "Bad Request",
				Message: "Invalid activity_group_id",
			})
		}

		response, err := controller.service.GetByActivityGroupID(activityGroupID)
		if err != nil {
			return ctx.Status(500).JSON(model.Response{
				Status:  "Server Error",
				Message: "Internal Server Error",
			})
		}

		return ctx.Status(200).JSON(model.Response{
			Status:  "Success",
			Message: "Success",
			Data:    response,
		})
	}

	todos, err := controller.service.GetAllTodoService()
	if err != nil {
		return ctx.Status(500).JSON(model.Response{
			Status:  "Server Error",
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func (controller *todoItemController) UpdateTodoItem(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	itemID, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(400).JSON(model.Response{
			Status:  "Bad Request",
			Message: "Invalid ID",
		})
	}

	request := new(model.TodoRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(model.Response{
			Status:  "Bad Request",
			Message: "Invalid request body",
		})
	}

	response, err := controller.service.UpdateTodoItem(itemID, *request)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", itemID),
			})
		}
		return ctx.Status(500).JSON(model.Response{
			Status:  "Server Error",
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Success",
		Message: "Success",
		Data:    response,
	})
}

func (controller *todoItemController) DeleteTodoController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "Invalid ID",
		})
	}
	err = controller.service.DeleteTodoService(uint(todoID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", todoID),
			})
		}
	}

	// Berikan respons sukses jika data berhasil dihapus
	return ctx.Status(200).JSON(model.Response{
		Status:  "Success",
		Message: "Success",
		Data:    fiber.Map{},
	})
}
