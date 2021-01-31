package configuration

import (
	"github.com/gruz0/monitoring-configuration-service/internal/persistence"
	"gorm.io/datatypes"
)

type Service interface {
	Configurations() (Configuration, error)
}

type service struct {
	persistence persistence.Persistence
}

func (s *service) Configurations() (Configuration, error) {
	var result Configuration

	query := "SELECT s.id AS site_id, s.domain_name, p.id AS plugin_id, p.namespace, p.name AS plugin_name, sp.settings " +
		"FROM sites s " +
		"INNER JOIN site_plugins sp ON sp.site_id = s.id " +
		"INNER JOIN plugins p ON sp.plugin_id = p.id " +
		"WHERE s.ownership_verified = true AND sp.enabled = true " +
		"ORDER BY s.id ASC, s.domain_name ASC, p.namespace ASC, p.name ASC"

	rows, err := s.persistence.DB.Raw(query).Rows()

	defer rows.Close()

	if err != nil {
		return result, err
	}

	var domains []Domain

	for rows.Next() {
		var (
			siteId          int
			domainName      string
			pluginId        int
			pluginNamespace string
			pluginName      string
			settings        datatypes.JSON
		)

		rows.Scan(&siteId, &domainName, &pluginId, &pluginNamespace, &pluginName, &settings)

		var domain *Domain
		domainFound := false
		domainIdx := -1

		for idx, d := range domains {
			if d.SiteID == siteId {
				domain = &d
				domainFound = true
				domainIdx = idx

				break
			}
		}

		if !domainFound {
			domain = &Domain{SiteID: siteId, Name: domainName, Plugins: []Plugin{}}
		}

		plugin := Plugin{ID: pluginId, Namespace: pluginNamespace, Name: pluginName, Settings: settings}

		domain.Plugins = append(domain.Plugins, plugin)

		if domainIdx >= 0 {
			domains[domainIdx] = *domain
		} else {
			domains = append(domains, *domain)
		}
	}

	result = Configuration{
		Domains: domains,
	}

	return result, nil
}

func NewService(p persistence.Persistence) Service {
	return &service{
		persistence: p,
	}
}
