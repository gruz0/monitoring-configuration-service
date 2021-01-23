package configuration

import (
	"github.com/gruz0/monitoring-configuration-service/internal/customer"
	"github.com/gruz0/monitoring-configuration-service/internal/plugin"
	"github.com/gruz0/monitoring-configuration-service/internal/site"
)

type Service interface {
	Configurations() []Configuration
}

type service struct {
	customers customer.Repository
}

func (s *service) Configurations() []Configuration {
	result := make([]Configuration, 0)

	for _, c := range s.customers.FindAllActive() {
		result = append(result, build(c))
	}

	return result
}

func NewService(customers customer.Repository) Service {
	return &service{
		customers: customers,
	}
}

type Configuration struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func build(c *customer.Customer) Configuration {
	domains := make([]Domain, 0)

	for _, s := range c.Sites {
		plugins := make([]Plugin, 0)

		for _, p := range s.Plugins {
			plugins = append(plugins, buildPlugin(p))
		}

		domains = append(domains, buildDomain(s, plugins))
	}

	return Configuration{Domains: domains}
}

func buildDomain(s *site.Site, plugins []Plugin) Domain {
	return Domain{
		ID:      s.Domain.ID,
		Name:    s.Domain.Name,
		Plugins: plugins,
	}
}

func buildPlugin(p *plugin.Plugin) Plugin {
	return Plugin{
		ID:   p.ID,
		Name: p.Name,
	}
}
