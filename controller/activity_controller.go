package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	CreateData(ctx *fiber.Ctx) error
	GetOneActivityController(ctx *fiber.Ctx) error
	UpdateActivity(ctx *fiber.Ctx) error
	Route(group fiber.Router)
}
