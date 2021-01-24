package plugin

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewPluginRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
