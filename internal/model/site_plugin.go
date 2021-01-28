package model

import "gorm.io/datatypes"

type SitePlugin struct {
	SiteID   int `gorm:"primaryKey;"`
	PluginID int `gorm:"primaryKey;"`

	Enabled bool `gorm:"index;default:true;"`

	Settings datatypes.JSON `gorm:"not null;default:'{}'::jsonb;"`
}
