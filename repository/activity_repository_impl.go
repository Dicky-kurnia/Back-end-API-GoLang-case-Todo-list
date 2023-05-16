package repository

import (
	"todo-list/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryActivity {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateActivity(input entity.Activity) (entity.Activity, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}
	return input, nil
}

func (r *repository) GetOneActivity(ID int) (entity.Activity, error) {
	var activity entity.Activity
	err := r.db.
		Where("id = ?", ID).
		Find(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *repository) GetAllActivity() ([]entity.Activity, error) {
	var activitys []entity.Activity
	err := r.db.Find(&activitys).Error
	if err != nil {
		return activitys, err
	}
	return activitys, nil
}

func (r *repository) UpdateActivity(activity entity.Activity) (entity.Activity, error) {
	err := r.db.Save(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *repository) DeleteActivity(id uint) error {
	var data entity.Activity
	if err := r.db.Where("id = ?", id).
		Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
