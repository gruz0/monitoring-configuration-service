package site

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewSiteRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
