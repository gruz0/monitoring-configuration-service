package site

import (
	"github.com/gruz0/monitoring-configuration-service/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) FindAllVerifiedDomainsWithPlugins() ([]model.Site, error) {
	var sites []model.Site

	err := r.db.Where(&model.Site{OwnershipVerified: true}).Preload("Plugins").Find(&sites).Error

	if err != nil {
		return nil, err
	}

	return sites, nil
}

func NewSiteRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
