package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Firstname string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Age       uint      `gorm:"not null"`
	Created   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
