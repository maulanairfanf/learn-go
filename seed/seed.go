package main

import (
	"log"
	"myapi/db"
	"myapi/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
    db.Init()

    // Seed users
    seedUsers()

    // Seed categories
    seedCategories()

    // Seed products
    seedProducts()

    log.Println("Seed data created successfully")
}

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to generate password hash: %v", err)
	}
	return string(hash)
}

func seedUsers() {
	// Create users
	users := []models.User{
		{
			Username:  "user1",
			Password:  hashPassword("password1"),
			Email:     "user1@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Username:  "user2",
			Password:  hashPassword("password2"),
			Email:     "user2@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		// Add more users as needed
	}

	for _, user := range users {
		if err := db.DB.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
	}
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
