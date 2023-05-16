package repository

import "todo-list/entity"

type RepositoryActivity interface {
	CreateActivity(data entity.Activity) (entity.Activity, error)
	GetOneActivity(ID int) (entity.Activity, error)
	GetAllActivity() ([]entity.Activity, error)
	UpdateActivity(activity entity.Activity) (entity.Activity, error)
	DeleteActivity(id uint) error
}
