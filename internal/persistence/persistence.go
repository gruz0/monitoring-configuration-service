package persistence

import (
	"github.com/gruz0/monitoring-configuration-service/internal/customer"
	"github.com/gruz0/monitoring-configuration-service/internal/model"
	"github.com/gruz0/monitoring-configuration-service/internal/plugin"
	"github.com/gruz0/monitoring-configuration-service/internal/site"
	"gorm.io/datatypes"
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

	if err := p.DB.Create(&plugins).Error; err != nil {
		return err
	}

	var customers = []model.Customer{
		{Email: "first@domain1.tld"},
		{Email: "second@domain2.tld"},
	}

	if err := p.DB.Create(&customers).Error; err != nil {
		return err
	}

	var sites = []model.Site{
		{DomainName: "domain1.tld", OwnershipVerified: true, CustomerID: customers[0].ID},
		{DomainName: "domain2.tld", OwnershipVerified: true, CustomerID: customers[1].ID},
		{DomainName: "domain3.tld", OwnershipVerified: false, CustomerID: customers[0].ID},
	}

	if err := p.DB.Create(&sites).Error; err != nil {
		return err
	}

	var sitePlugins = []model.SitePlugin{
		{SiteID: sites[0].ID, PluginID: plugins[0].ID, Settings: datatypes.JSON([]byte(`{"resource":"/resource1","value":"content1"}`))}, // content.contains_string
		{SiteID: sites[0].ID, PluginID: plugins[1].ID, Settings: datatypes.JSON([]byte(`{"resource":"/resource2","value":"content2"}`))}, // content.does_not_contain_string
		{SiteID: sites[0].ID, PluginID: plugins[2].ID},  // content.valid_json
		{SiteID: sites[1].ID, PluginID: plugins[3].ID},  // domain.dns_a_records
		{SiteID: sites[1].ID, PluginID: plugins[4].ID},  // domain.domain_expiration
		{SiteID: sites[1].ID, PluginID: plugins[5].ID},  // domain.ssl_certificate_expiration
		{SiteID: sites[2].ID, PluginID: plugins[6].ID},  // files.directory_listing_disabled
		{SiteID: sites[2].ID, PluginID: plugins[7].ID},  // files.file_does_not_exist
		{SiteID: sites[2].ID, PluginID: plugins[8].ID},  // files.file_exists
		{SiteID: sites[0].ID, PluginID: plugins[9].ID},  // files.robots_txt
		{SiteID: sites[0].ID, PluginID: plugins[10].ID}, // files.sitemap_xml
		{SiteID: sites[0].ID, PluginID: plugins[11].ID}, // http.http_status200
		{SiteID: sites[0].ID, PluginID: plugins[12].ID}, // http.http_to_https_redirect
		{SiteID: sites[0].ID, PluginID: plugins[13].ID}, // http.non_existent_url_returns404
		{SiteID: sites[0].ID, PluginID: plugins[14].ID, Settings: datatypes.JSON([]byte(`{"resource":"/resource","value":301}`))}, // http.valid_http_status_code
		{SiteID: sites[0].ID, PluginID: plugins[15].ID}, // http.www_to_non_www_redirect
		{SiteID: sites[0].ID, PluginID: plugins[16].ID}, // other.database_connection_issue
		{SiteID: sites[0].ID, PluginID: plugins[17].ID}, // ownership.dns_a_records
		{SiteID: sites[0].ID, PluginID: plugins[18].ID}, // ownership.file_verification
	}

	if err := p.DB.Create(&sitePlugins).Error; err != nil {
		return err
	}

	return nil
}

func (p *Persistence) DeleteAll() error {
	if err := p.DB.Exec("DELETE FROM site_plugins;").Error; err != nil {
		return err
	}

	if err := p.DB.Exec("DELETE FROM plugins;").Error; err != nil {
		return err
	}

	if err := p.DB.Exec("DELETE FROM sites;").Error; err != nil {
		return err
	}

	if err := p.DB.Exec("DELETE FROM customers;").Error; err != nil {
		return err
	}

	return nil
}
