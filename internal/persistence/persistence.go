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

	// NOTE: https://gorm.io/docs/many_to_many.html#Customize-JoinTable
	err = db.SetupJoinTable(&model.Site{}, "Plugins", &model.SitePlugin{})

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
		&model.SitePlugin{},
	)
}

func (p *Persistence) Seed() error {
	var plugins = []model.Plugin{
		// Content
		{Namespace: "content", Name: "contains_string", Title: "Contains String", Description: "Description"},
		{Namespace: "content", Name: "does_not_contain_string", Title: "Does Not Contain String", Description: "Description"},
		{Namespace: "content", Name: "valid_json", Title: "Valid JSON", Description: "Description"},

		// Domain
		{Namespace: "domain", Name: "dns_a_records", Title: "DNS A Records", Description: "Description"},
		{Namespace: "domain", Name: "domain_expiration", Title: "Domain Expiration", Description: "Description"},
		{Namespace: "domain", Name: "ssl_certificate_expiration", Title: "SSL Expiration", Description: "Description"},

		// Files
		{Namespace: "files", Name: "directory_listing_disabled", Title: "Directory Listing Disabled", Description: "Description"},
		{Namespace: "files", Name: "file_does_not_exist", Title: "File Does Not Exist", Description: "Description"},
		{Namespace: "files", Name: "file_exists", Title: "File Exists", Description: "Description"},
		{Namespace: "files", Name: "robots_txt", Title: "robots.txt", Description: "Description"},
		{Namespace: "files", Name: "sitemap_xml", Title: "sitemap.xml", Description: "Description"},

		// HTTP
		{Namespace: "http", Name: "http_status200", Title: "HTTP Status 200", Description: "Description"},
		{Namespace: "http", Name: "http_to_https_redirect", Title: "HTTP to HTTPS Redirect", Description: "Description"},
		{Namespace: "http", Name: "non_existent_url_returns404", Title: "Non Existent URL Returns HTTP Status 404", Description: "Description"},
		{Namespace: "http", Name: "valid_http_status_code", Title: "Valid HTTP Status", Description: "Description"},
		{Namespace: "http", Name: "www_to_non_www_redirect", Title: "WWW to non-WWW Redirect", Description: "Description"},

		// Other
		{Namespace: "other", Name: "database_connection_issue", Title: "Database Connection Issue", Description: "Description"},

		// Ownership
		{Namespace: "ownership", Name: "dns_txt_record_verification", Title: "DNS TXT Record Verification", Description: "Description"},
		{Namespace: "ownership", Name: "file_verification", Title: "File Verification", Description: "Description"},

		// WordPress
	}

	p.DB.Create(&plugins)

	var customers = []model.Customer{
		{Email: "first@domain1.tld"},
		{Email: "second@domain2.tld"},
	}

	p.DB.Create(&customers)

	var sites = []model.Site{
		{DomainName: "domain1.tld", OwnershipVerified: true, CustomerID: customers[0].ID},
		{DomainName: "domain2.tld", OwnershipVerified: true, CustomerID: customers[1].ID},
		{DomainName: "domain3.tld", OwnershipVerified: false, CustomerID: customers[0].ID},
	}

	p.DB.Create(&sites)

	var sitePlugins = []model.SitePlugin{
		{SiteID: sites[0].ID, PluginID: plugins[0].ID},
		{SiteID: sites[0].ID, PluginID: plugins[1].ID},
		{SiteID: sites[0].ID, PluginID: plugins[2].ID},
		{SiteID: sites[1].ID, PluginID: plugins[3].ID},
		{SiteID: sites[1].ID, PluginID: plugins[4].ID},
		{SiteID: sites[1].ID, PluginID: plugins[5].ID},
		{SiteID: sites[2].ID, PluginID: plugins[6].ID},
		{SiteID: sites[2].ID, PluginID: plugins[7].ID},
		{SiteID: sites[2].ID, PluginID: plugins[8].ID},
	}

	p.DB.Create(&sitePlugins)

	return nil
}
