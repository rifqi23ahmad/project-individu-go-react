package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	// Role      []Role `gorm:"foreignKey:User"`
}

type Role struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
}

type Student struct {
	ID int `gorm:"primaryKey"`
	// Mentor  []Mentor  `gorm:"foreignKey:MentorID"`
	// Booking []Booking `gorm:"foreignKey:BookingID"`
}

type Mentor struct {
	ID int `gorm:"primaryKey"`
	// User    []User    `gorm:"foreignKey:UserID"`
	// Booking []Booking `gorm:"foreignKey:BookingID"`
}

type Booking struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	Time        time.Time
	Date        time.Time
	IsCompleted bool
}
