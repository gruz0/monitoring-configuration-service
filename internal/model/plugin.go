package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Plugin struct {
	gorm.Model

	ID        int
	Namespace string `gorm:"index;not null;"`
	Name      string `gorm:"uniqueIndex;not null;"`

	Title       string `gorm:"not null;"`
	Description string `gorm:"not null;"`

	Settings datatypes.JSON
}
