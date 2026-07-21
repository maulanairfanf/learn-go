package models

import "time"

type Product struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"type:varchar(100);not null" binding:"required,min=3"`
	Quantity    int        `json:"quantity" gorm:"not null"`
	Categories  []Category `json:"categories" gorm:"many2many:product_categories;"`
	Price       float64    `json:"price" gorm:"not null"`
	Description string     `json:"description" gorm:"type:text"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Categories  []uint  `json:"categories"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
