package configuration

import (
	"github.com/gruz0/monitoring-configuration-service/internal/model"
	"github.com/gruz0/monitoring-configuration-service/internal/site"
)

type Service interface {
	Configurations() []Configuration
}

type service struct {
	sites site.Repository
}

func (s *service) Configurations() []Configuration {
	sites, err := s.sites.FindAllVerifiedDomainsWithPlugins()

	if err != nil {
		return nil
	}

	return []Configuration{
		{
			Domains: buildDomains(sites),
		},
	}
}

func NewService(sites site.Repository) Service {
	return &service{
		sites: sites,
	}
}

func buildDomains(sites []model.Site) []Domain {
	domains := make([]Domain, 0)

	for _, site := range sites {
		if len(site.Plugins) == 0 {
			continue
		}

		domains = append(
			domains,
			Domain{
				SiteID:  site.ID,
				Name:    site.DomainName,
				Plugins: buildPlugins(site.Plugins),
			},
		)
	}

	return domains
}

func buildPlugins(plugins []model.Plugin) []Plugin {
	p := make([]Plugin, len(plugins))

	for i, plugin := range plugins {
		p[i] = Plugin{
			ID:        plugin.ID,
			Namespace: plugin.Namespace,
			Name:      plugin.Name,
		}
	}

	return p
}
