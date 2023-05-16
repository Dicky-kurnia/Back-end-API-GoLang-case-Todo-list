package controller

import "github.com/gofiber/fiber/v2"

type TodoItemController interface {
	CreateTodoItem(ctx *fiber.Ctx) error
	GetTodoItem(ctx *fiber.Ctx) error
	GetActivityGroupIdController(ctx *fiber.Ctx) error
	GetAllTodoController(ctx *fiber.Ctx) error
	UpdateTodoItem(ctx *fiber.Ctx) error
	DeleteTodoController(ctx *fiber.Ctx) error
	Route(group fiber.Router)
}
