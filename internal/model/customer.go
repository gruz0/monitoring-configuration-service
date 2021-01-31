package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();"`
	Email string    `gorm:"uniqueIndex;not null;"`
	Sites []Site    `gorm:"constraint:OnDelete:CASCADE;"`
}
