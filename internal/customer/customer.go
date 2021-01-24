package customer

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
