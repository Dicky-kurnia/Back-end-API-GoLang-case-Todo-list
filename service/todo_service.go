package service

import (
	"Back-end-API-GoLang-case-Todo-list/entity"
	"Back-end-API-GoLang-case-Todo-list/model"
)

type TodoItemService interface {
	Create(itemRequest model.TodoRequest) (*model.TodoResponse, error)
	GetOneTodoService(itemID int) (*model.TodoResponse, error)
	GetByActivityGroupID(activityGroupID int) ([]model.TodoResponse, error)
	UpdateTodoItem(itemID int, request model.TodoRequest) (entity.Todo, error)
	DeleteTodoService(id uint) error
	GetAllTodoService() ([]entity.Todo, error)
}
