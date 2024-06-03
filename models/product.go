package models

import "time"

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Quantity    int        `gorm:"not null" json:"quantity"`
	Categories  []Category `gorm:"many2many:product_categories;" json:"categories"` // Many-to-many relationship
	Price       string     `gorm:"type:varchar(100);not null" json:"price"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}