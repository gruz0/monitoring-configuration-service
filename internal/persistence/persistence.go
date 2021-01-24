package persistence

import (
	"github.com/gruz0/monitoring-configuration-service/internal/customer"
	"github.com/gruz0/monitoring-configuration-service/internal/model"
	"github.com/gruz0/monitoring-configuration-service/internal/plugin"
	"github.com/gruz0/monitoring-configuration-service/internal/site"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Persistence struct {
	DB        *gorm.DB
	Customers *customer.Repository
	Sites     *site.Repository
	Plugins   *plugin.Repository
}

func New(dsn string) (*Persistence, error) {
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)

	if err != nil {
		return nil, err
	}

	return &Persistence{
		DB:        db,
		Customers: newCustomerRepository(db),
		Sites:     newSiteRepository(db),
		Plugins:   newPluginRepository(db),
	}, nil
}

func newCustomerRepository(db *gorm.DB) *customer.Repository {
	return customer.NewCustomerRepository(db)
}

func newSiteRepository(db *gorm.DB) *site.Repository {
	return site.NewSiteRepository(db)
}

func newPluginRepository(db *gorm.DB) *plugin.Repository {
	return plugin.NewPluginRepository(db)
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Customer{},
		&model.Site{},
		&model.Plugin{},
	)
}
