package model

type SitePlugin struct {
	SiteID   int `gorm:"primaryKey;"`
	PluginID int `gorm:"primaryKey;"`

	// FIXME: Add test case to check that FindAllVerifiedDomainsWithPlugins returns only enabled plugins
	Enabled bool `gorm:"index;default:true;"`
}
