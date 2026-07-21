package repositories

import (
	"myapi/db"
	"myapi/models"
)

type CategoryRepository struct{}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := db.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id int) (*models.Category, error) {
	var category models.Category
	err := db.DB.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return db.DB.Create(category).Error
}

func (r *CategoryRepository) Save(category *models.Category) error {
	return db.DB.Save(category).Error
}

func (r *CategoryRepository) Delete(category *models.Category) error {
	return db.DB.Delete(category).Error
}
