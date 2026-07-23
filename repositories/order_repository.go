package repositories

import (
	"myapi/db"
	"myapi/models"

	"gorm.io/gorm"
)

type OrderRepository struct{}

func (r *OrderRepository) FindAll() ([]models.Order, error) {
	var order []models.Order
	err := db.DB.Preload("Items.Product.Categories").Preload("Items.Product").Find(&order).Error
	return order, err
}

func (r *OrderRepository) FindByID(id int) (*models.Order, error) {
	var order models.Order
	err := db.DB.Preload("Items.Product.Categories").Preload("Items.Product").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) Create(order *models.Order) (*models.Order, error) {
	err := db.DB.Create(&order).Error
	return order, err
}

func (r *OrderRepository) Save(order *models.Order) (*models.Order, error) {
	err := db.DB.Save(&order).Error
	return order, err
}

func (r *OrderRepository) DecreaseStock(tx *gorm.DB, productID uint, qty int) error {
	return tx.Model(&models.Product{}).Where("id = ?", productID).
		Update("quantity", gorm.Expr("quantity - ?", qty)).Error
}

func (r *OrderRepository) IncreaseStock(tx *gorm.DB, productID uint, qty int) error {
	return tx.Model(&models.Product{}).Where("id = ?", productID).
		Update("quantity", gorm.Expr("quantity + ?", qty)).Error
}

func (r *OrderRepository) DeleteOrder(tx *gorm.DB, id uint) error {
	tx.Where("order_id = ?", id).Delete(&models.OrderItem{})
	return tx.Delete(&models.Order{}, id).Error
}
