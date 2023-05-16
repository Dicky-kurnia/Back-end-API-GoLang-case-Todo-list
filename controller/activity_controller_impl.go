package controller

import (
	"Back-end-API-GoLang-case-Todo-list/entity"
	"Back-end-API-GoLang-case-Todo-list/model"
	"Back-end-API-GoLang-case-Todo-list/service"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type controller struct {
	service service.ServiceActivity
}

func NewController(service service.ServiceActivity) Controller {
	return &controller{
		service: service,
	}
}

func (controller *controller) Route(group fiber.Router) {
	group.Post("", controller.CreateData)
	group.Get("/:id", controller.GetOneActivityController)
	group.Get("", controller.GetAllActivityController)
	group.Patch("/:id", controller.UpdateActivity)
	group.Delete("/:id", controller.DeleteDataController)

}

func (controller *controller) CreateData(ctx *fiber.Ctx) error {
	var data entity.Activity
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	newActivity, err := controller.service.CreateData(data)
	if err != nil {
		return ctx.Status(400).JSON(model.Response{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Success",
		Message: "Success",
		Data:    newActivity,
	})
}

func (controller *controller) GetOneActivityController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data := entity.Activity{}

	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "Invalid ID",
		})
	}
	data.ID = parsedID

	activity, err := controller.service.GetOneActivityService(data)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %d Not Found", parsedID),
			})
		}

	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Succsess",
		Message: "Succsess",
		Data:    activity,
	})
}

func (controller *controller) GetAllActivityController(ctx *fiber.Ctx) error {
	activities, err := controller.service.GetAllActivityService()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: "Activity Not Found",
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
		Data:    activities,
	})
}

func (controller *controller) DeleteDataController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	dataID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "Invalid ID",
		})
	}
	err = controller.service.DeleteActivityService(uint(dataID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %d Not Found", dataID),
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

func (controller *controller) UpdateActivity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(model.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %d Not Found", idInt),
			})
		}
	}

	var inputID model.ActivityInputID
	inputID.ID = idInt

	var inputData model.ActivityInput
	if err := ctx.BodyParser(&inputData); err != nil {
		return ctx.Status(400).JSON(model.Response{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
	}

	updatedActivity, err := controller.service.UpdateActivity(inputID, inputData)
	if err != nil {
		return ctx.Status(500).JSON(model.Response{
			Status: "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(model.Response{
		Status:  "Succsess",
		Message: "Succsess",
		Data:    updatedActivity,
	})
}
