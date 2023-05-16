package service

import (
	"Back-end-API-GoLang-case-Todo-list/entity"
	"Back-end-API-GoLang-case-Todo-list/model"
	"Back-end-API-GoLang-case-Todo-list/repository"
)

type service struct {
	repo repository.RepositoryActivity
}

func NewService(repo repository.RepositoryActivity) ServiceActivity {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateData(input entity.Activity) (entity.Activity, error) {
	activity := entity.Activity{}
	activity.Title = input.Title
	activity.Email = input.Email

	newActivity, err := s.repo.CreateActivity(input)
	if err != nil {
		return newActivity, err
	}
	return newActivity, nil
}

func (s *service) GetOneActivityService(input entity.Activity) (entity.Activity, error) {
	activity, err := s.repo.GetOneActivity(input.ID)
	if err != nil {
		return entity.Activity{}, err
	}
	return activity, nil
}

func (s *service) GetAllActivityService() ([]entity.Activity, error) {
	activitys, err := s.repo.GetAllActivity()
	if err != nil {
		return nil, err
	}
	return activitys, nil
}

func (s *service) DeleteActivityService(id uint) error {
	err := s.repo.DeleteActivity(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateActivity(inputID model.ActivityInputID, inputData model.ActivityInput) (entity.Activity, error) {
	activity, err := s.repo.GetOneActivity(inputID.ID)
	if err != nil {
		return activity, err
	}

	activity.Title = inputData.Title

	updatedActivity, err := s.repo.UpdateActivity(activity)
	if err != nil {
		return updatedActivity, err
	}

	return updatedActivity, nil
}
