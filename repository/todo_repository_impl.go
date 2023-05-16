package repository

import (
	"Back-end-API-GoLang-case-Todo-list/entity"

	"gorm.io/gorm"
)

type todoItemRepository struct {
	db *gorm.DB
}

func NewRepositoryTod(db *gorm.DB) TodoItemRepository {
	return &todoItemRepository{
		db: db,
	}
}

func (repository *todoItemRepository) Create(item entity.Todo) (entity.Todo, error) {
	err := repository.db.Create(&item).Error
	if err != nil {
		return item, err
	}
	return item, nil
}

func (repository *todoItemRepository) GetOneTodo(itemID int) (*entity.Todo, error) {
	item := new(entity.Todo)
	result := repository.db.Where("id = ?", itemID).Find(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (repository *todoItemRepository) GetByActivityGroupID(activityGroupID int) ([]entity.Todo, error) {
	var items []entity.Todo
	result := repository.db.
		Where("activity_group_id = ?", activityGroupID).
		Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (repository *todoItemRepository) GetAllTodo() ([]entity.Todo, error) {
	var todos []entity.Todo
	err := repository.db.Find(&todos).Error
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (repository *todoItemRepository) UpdateTodoItem(todo entity.Todo) (entity.Todo, error) {
	err := repository.db.Save(&todo).Error
	if err != nil {
		return entity.Todo{}, err
	}
	return todo, nil
}

func (repository *todoItemRepository) DeleteTodoItem(id uint) error {
	var data entity.Todo
	if err := repository.db.Where("id = ?", id).
		Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
