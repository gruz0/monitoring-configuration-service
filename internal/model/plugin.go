package model

import "gorm.io/gorm"

type Plugin struct {
	gorm.Model

	ID   int
	Name string `gorm:"uniqueIndex;not null;"`

	Sites []Site `gorm:"many2many:site_plugins;constraint:OnDelete:CASCADE;"`
}
