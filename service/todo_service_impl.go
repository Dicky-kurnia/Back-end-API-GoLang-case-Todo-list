package service

import (
	"todo-list/entity"
	"todo-list/model"
	"todo-list/repository"
)

type todoItemService struct {
	repository repository.TodoItemRepository
}

func NewTodoItemServiceTodo(repository repository.TodoItemRepository) TodoItemService {
	return &todoItemService{
		repository: repository,
	}
}

func (service *todoItemService) Create(itemRequest model.TodoRequest) (*model.TodoResponse, error) {
	item := &entity.Todo{
		Title:           itemRequest.Title,
		ActivityGroupID: itemRequest.ActivityGroupID,
		IsActive:        itemRequest.IsActive,
		Priority:        itemRequest.Priority,
	}

	createdItem, err := service.repository.Create(*item)
	if err != nil {
		return nil, err
	}

	response := &model.TodoResponse{
		ID:              createdItem.ID,
		Title:           createdItem.Title,
		ActivityGroupID: createdItem.ActivityGroupID,
		IsActive:        createdItem.IsActive,
		Priority:        createdItem.Priority,
		UpdatedAt:       createdItem.UpdatedAt.String(),
		CreatedAt:       createdItem.CreatedAt.String(),
	}

	return response, nil
}

func (service *todoItemService) GetOneTodoService(itemID int) (*model.TodoResponse, error) {
	item, err := service.repository.GetOneTodo(itemID)
	if err != nil {
		return nil, err
	}

	response := &model.TodoResponse{
		ID:              item.ID,
		Title:           item.Title,
		ActivityGroupID: item.ActivityGroupID,
		IsActive:        item.IsActive,
		Priority:        item.Priority,
		UpdatedAt:       item.UpdatedAt.String(),
		CreatedAt:       item.CreatedAt.String(),
	}

	return response, nil
}

func (service *todoItemService) GetByActivityGroupID(activityGroupID int) ([]model.TodoResponse, error) {
	items, err := service.repository.GetByActivityGroupID(activityGroupID)
	if err != nil {
		return nil, err
	}

	response := make([]model.TodoResponse, len(items))
	for i, item := range items {
		response[i] = model.TodoResponse{
			ID:              item.ID,
			Title:           item.Title,
			ActivityGroupID: item.ActivityGroupID,
			IsActive:        item.IsActive,
			Priority:        item.Priority,
			UpdatedAt:       item.UpdatedAt.String(),
			CreatedAt:       item.CreatedAt.String(),
		}
	}

	return response, nil
}

func (service *todoItemService) GetAllTodoService() ([]entity.Todo, error) {
	activitys, err := service.repository.GetAllTodo()
	if err != nil {
		return nil, err
	}
	return activitys, nil
}

func (service *todoItemService) UpdateTodoItem(itemID int, request model.TodoRequest) (entity.Todo, error) {
	todo, err := service.repository.GetOneTodo(itemID)
	if err != nil {
		return entity.Todo{}, err
	}

	todo.Title = request.Title
	todo.IsActive = request.IsActive
	todo.Priority = request.Priority

	updatedTodo, err := service.repository.UpdateTodoItem(*todo)
	if err != nil {
		return entity.Todo{}, err
	}

	return updatedTodo, nil
}
func (service *todoItemService) DeleteTodoService(id uint) error {
	err := service.repository.DeleteTodoItem(id)
	if err != nil {
		return err
	}
	return nil
}
