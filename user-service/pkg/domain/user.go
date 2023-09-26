package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
