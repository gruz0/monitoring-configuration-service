package persistence

import (
	"github.com/gruz0/monitoring-configuration-service/internal/customer"
	"github.com/gruz0/monitoring-configuration-service/internal/domain"
	"github.com/gruz0/monitoring-configuration-service/internal/plugin"
	"github.com/gruz0/monitoring-configuration-service/internal/site"
)

func NewCustomerRepository() *customer.Repository {
	return customer.NewCustomerRepository()
}

func NewSiteRepository() *site.Repository {
	return site.NewSiteRepository()
}

func NewDomainRepository() *domain.Repository {
	return domain.NewDomainRepository()
}

func NewPluginRepository() *plugin.Repository {
	return plugin.NewPluginRepository()
}
