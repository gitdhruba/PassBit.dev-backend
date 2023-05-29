package models

//This package contains all model definitions
//Author : Dhruba Sinha

//this file specifically contains model definitions related to users and authentication

import (
	"time"

	"gorm.io/gorm"
)

// GenerateISOString generates a time string equivalent to Date.now().toISOString in JavaScript
func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// Base struct for User model
// - this is to be used only by gorm for unique identification for each user
type Base struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt string
	UpdatedAt string
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	// generate timestamps
	t := GenerateISOString()
	b.CreatedAt, b.UpdatedAt = t, t
	return nil
}

func (b *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	b.UpdatedAt = GenerateISOString()
	return nil
}

// User model definition
type User struct {
	Base
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}
