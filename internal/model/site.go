package model

import "gorm.io/gorm"

type Site struct {
	gorm.Model

	ID                int
	DomainName        string   `gorm:"uniqueIndex;not null;"`
	OwnershipVerified bool     `gorm:"not null;default:false;"`
	Plugins           []Plugin `gorm:"many2many:site_plugins;association_autocreate:false;constraint:OnDelete:CASCADE;"`

	CustomerID int
}
