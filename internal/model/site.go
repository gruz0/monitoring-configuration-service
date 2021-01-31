package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model

	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();"`
	DomainName        string    `gorm:"uniqueIndex;not null;"`
	OwnershipVerified bool      `gorm:"not null;default:false;"`
	Plugins           []Plugin  `gorm:"many2many:site_plugins;association_autocreate:false;constraint:OnDelete:CASCADE;"`

	CustomerID uuid.UUID `gorm:"type:uuid;"`
}
