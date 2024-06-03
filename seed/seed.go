package main

import (
	"log"
	"myapi/db"
	"myapi/models"
	"time"
)

func main() {
    db.Init()

    // Seed categories
    seedCategories()

    // Seed products
    seedProducts()

    log.Println("Seed data created successfully")
}

func seedCategories() {
    categories := []models.Category{
        {Name: "Category 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
        {Name: "Category 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
        // Add more categories as needed
    }

    for _, category := range categories {
        if err := db.DB.Create(&category).Error; err != nil {
            log.Fatalf("Failed to create category: %v", err)
        }
    }
}

func seedProducts() {
    // Retrieve categories
    var categories []models.Category
    if err := db.DB.Find(&categories).Error; err != nil {
        log.Fatalf("Failed to retrieve categories: %v", err)
    }

    products := []models.Product{
        {
            Name:        "Product 1",
            Quantity:    10,
            Categories:  []models.Category{categories[0]}, // Assign the first category to Product 1
            Price:       "200000",
            Description: "Description of Product 1",
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
            Name:        "Product 2",
            Quantity:    20,
            Categories:  []models.Category{categories[1]}, // Assign the second category to Product 2
            Price:       "1000000",
            Description: "Description of Product 2",
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        // Add more products as needed
    }

    for _, product := range products {
        if err := db.DB.Create(&product).Error; err != nil {
            log.Fatalf("Failed to create product: %v", err)
        }
    }
}
