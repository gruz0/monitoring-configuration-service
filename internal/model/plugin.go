package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Plugin struct {
	gorm.Model

	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();"`
	Namespace string    `gorm:"index;not null;"`
	Name      string    `gorm:"uniqueIndex;not null;"`

	Title       string `gorm:"not null;"`
	Description string `gorm:"not null;"`

	Settings datatypes.JSON
}
