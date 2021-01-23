package site

import (
	"github.com/gruz0/monitoring-configuration-service/internal/domain"
	"github.com/gruz0/monitoring-configuration-service/internal/plugin"
)

type Site struct {
	ID      int
	Domain  *domain.Domain
	Plugins []*plugin.Plugin
}

type Repository struct{}

func New(id int, domain *domain.Domain, plugins []*plugin.Plugin) *Site {
	return &Site{
		ID:      id,
		Domain:  domain,
		Plugins: plugins,
	}
}

func (r *Repository) FindAll() []*Site {
	return []*Site{}
}

func (r *Repository) FindAllByCustomerID(_ int) []*Site {
	return []*Site{}
}

func NewSiteRepository() *Repository {
	return &Repository{}
}
