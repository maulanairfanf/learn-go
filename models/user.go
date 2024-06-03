package models

import "time"

// User represents a user account in the database
type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"type:varchar(50);unique;not null" json:"username"`
    Password  string    `gorm:"type:varchar(100);not null" json:"-"`
    Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
