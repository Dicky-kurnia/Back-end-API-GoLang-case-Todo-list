package service

import (
	"Back-end-API-GoLang-case-Todo-list/entity"
	"Back-end-API-GoLang-case-Todo-list/model"
)

type ServiceActivity interface {
	CreateData(input entity.Activity) (entity.Activity, error)
	GetOneActivityService(input entity.Activity) (entity.Activity, error)
	GetAllActivityService() ([]entity.Activity, error)
	UpdateActivity(inputID model.ActivityInputID, inputData model.ActivityInput) (entity.Activity, error)
	DeleteActivityService(id uint) error
}
