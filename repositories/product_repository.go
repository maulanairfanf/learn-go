package repositories

import (
	"myapi/db"
	"myapi/models"
)

type ProductRepository struct{}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Preload("Categories").Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id int) (*models.Product, error) {
	var product models.Product
	err := db.DB.Preload("Categories").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) FindCategoriesByIDs(ids []uint) ([]models.Category, error) {
	var categories []models.Category
	err := db.DB.Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}

func (r *ProductRepository) Create(product *models.Product) error {
	return db.DB.Create(product).Error
}

func (r *ProductRepository) ReplaceCategories(product *models.Product, categories []models.Category) error {
	return db.DB.Model(product).Association("Categories").Replace(categories)
}

func (r *ProductRepository) Save(product *models.Product) error {
	return db.DB.Save(product).Error
}

func (r *ProductRepository) ClearCategories(product *models.Product) error {
	return db.DB.Model(product).Association("Categories").Clear()
}

func (r *ProductRepository) Delete(product *models.Product) error {
	return db.DB.Delete(product).Error
}

func (r *ProductRepository) PreloadCategories(product *models.Product) error {
	return db.DB.Preload("Categories").First(product, product.ID).Error
}
