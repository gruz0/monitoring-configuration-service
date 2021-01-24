package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	ID    int
	Email string `gorm:"uniqueIndex;not null;"`
	Sites []Site `gorm:"constraint:OnDelete:CASCADE;"`
}
