package configuration

import (
	"testing"

	"github.com/gruz0/monitoring-configuration-service/internal/persistence"
	"github.com/stretchr/testify/assert"
)

func TestConfigurations(t *testing.T) {
	configurationDatabaseURL := "host=localhost user=app password=password dbname=app_development sslmode=disable TimeZone=UTC"

	db, err := persistence.New(configurationDatabaseURL)

	if err != nil {
		t.Fatalf("Unable to connect to a database: %s", err)
	}

	if err := persistence.AutoMigrate(db.DB); err != nil {
		t.Fatalf("Unable to migrate a database: %s", err)
	}

	if err := db.DeleteAll(); err != nil {
		t.Fatalf("Unable to delete records: %s", err)
	}

	if err := db.Seed(); err != nil {
		t.Fatalf("Unable to seed a database: %s", err)
	}

	cs := NewService(*db)
	c, err := cs.Configurations()

	if err != nil {
		t.Fatalf("Unable to fetch configurations: %s", err)
	}

	assert.Equal(t, 2, len(c.Domains))

	firstDomain := c.Domains[0]

	assert.Equal(t, 13, len(firstDomain.Plugins))

	assert.Equal(t, "content", firstDomain.Plugins[0].Namespace)
	assert.Equal(t, "contains_string", firstDomain.Plugins[0].Name)
	assert.Equal(t, `{"value": "content1", "resource": "/resource1"}`, string(firstDomain.Plugins[0].Settings))

	assert.Equal(t, "content", firstDomain.Plugins[1].Namespace)
	assert.Equal(t, "does_not_contain_string", firstDomain.Plugins[1].Name)
	assert.Equal(t, `{"value": "content2", "resource": "/resource2"}`, string(firstDomain.Plugins[1].Settings))

	assert.Equal(t, "content", firstDomain.Plugins[2].Namespace)
	assert.Equal(t, "valid_json", firstDomain.Plugins[2].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[2].Settings))

	assert.Equal(t, "files", firstDomain.Plugins[3].Namespace)
	assert.Equal(t, "robots_txt", firstDomain.Plugins[3].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[3].Settings))

	assert.Equal(t, "files", firstDomain.Plugins[4].Namespace)
	assert.Equal(t, "sitemap_xml", firstDomain.Plugins[4].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[4].Settings))

	assert.Equal(t, "http", firstDomain.Plugins[5].Namespace)
	assert.Equal(t, "http_status200", firstDomain.Plugins[5].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[5].Settings))

	assert.Equal(t, "http", firstDomain.Plugins[6].Namespace)
	assert.Equal(t, "http_to_https_redirect", firstDomain.Plugins[6].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[6].Settings))

	assert.Equal(t, "http", firstDomain.Plugins[7].Namespace)
	assert.Equal(t, "non_existent_url_returns404", firstDomain.Plugins[7].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[7].Settings))

	assert.Equal(t, "http", firstDomain.Plugins[8].Namespace)
	assert.Equal(t, "valid_http_status_code", firstDomain.Plugins[8].Name)
	assert.Equal(t, `{"value": 301, "resource": "/resource"}`, string(firstDomain.Plugins[8].Settings))

	assert.Equal(t, "http", firstDomain.Plugins[9].Namespace)
	assert.Equal(t, "www_to_non_www_redirect", firstDomain.Plugins[9].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[9].Settings))

	assert.Equal(t, "other", firstDomain.Plugins[10].Namespace)
	assert.Equal(t, "database_connection_issue", firstDomain.Plugins[10].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[10].Settings))

	assert.Equal(t, "ownership", firstDomain.Plugins[11].Namespace)
	assert.Equal(t, "dns_txt_record_verification", firstDomain.Plugins[11].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[11].Settings))

	assert.Equal(t, "ownership", firstDomain.Plugins[12].Namespace)
	assert.Equal(t, "file_verification", firstDomain.Plugins[12].Name)
	assert.Equal(t, `{}`, string(firstDomain.Plugins[12].Settings))

	secondDomain := c.Domains[1]

	assert.Equal(t, 3, len(secondDomain.Plugins))

	assert.Equal(t, "domain", secondDomain.Plugins[0].Namespace)
	assert.Equal(t, "dns_a_records", secondDomain.Plugins[0].Name)
	assert.Equal(t, `{}`, string(secondDomain.Plugins[0].Settings))

	assert.Equal(t, "domain", secondDomain.Plugins[1].Namespace)
	assert.Equal(t, "domain_expiration", secondDomain.Plugins[1].Name)
	assert.Equal(t, `{}`, string(secondDomain.Plugins[1].Settings))

	assert.Equal(t, "domain", secondDomain.Plugins[2].Namespace)
	assert.Equal(t, "ssl_certificate_expiration", secondDomain.Plugins[2].Name)
	assert.Equal(t, `{}`, string(secondDomain.Plugins[2].Settings))
}
