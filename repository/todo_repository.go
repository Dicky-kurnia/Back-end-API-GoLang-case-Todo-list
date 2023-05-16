package repository

import "todo-list/entity"

type TodoItemRepository interface {
	Create(item entity.Todo) (entity.Todo, error)
	GetOneTodo(itemID int) (*entity.Todo, error)
	GetByActivityGroupID(activityGroupID int) ([]entity.Todo, error)
	UpdateTodoItem(todo entity.Todo) (entity.Todo, error)
	DeleteTodoItem(id uint) error
	GetAllTodo() ([]entity.Todo, error)
}
