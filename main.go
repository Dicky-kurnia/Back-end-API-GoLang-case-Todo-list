package main

import (
	"log"
	"todo-list/config"
	"todo-list/controller"
	"todo-list/repository"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Buat instance repository
	repositoryActivity := repository.NewRepository(db)
	repositoryTodo := repository.NewRepositoryTod(db)

	// Buat instance service
	serviceActivity := service.NewService(repositoryActivity)
	todoService := service.NewTodoItemServiceTodo(repositoryTodo)

	// Buat instance controller
	controllerActivity := controller.NewController(serviceActivity)
	todoController := controller.NewControllerTodo(todoService)

	// activity group
	activity := app.Group("/activity-groups")

	// todo group
	todo := app.Group("/todo-items")

	// Implement Group Routes Here
	controllerActivity.Route(activity)
	todoController.Route(todo)

	// Jalankan server
	log.Fatal(app.Listen(":3030"))
}
