package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type SitePlugin struct {
	SiteID   uuid.UUID `gorm:"primaryKey;type:uuid;"`
	PluginID uuid.UUID `gorm:"primaryKey;type:uuid;"`

	Enabled bool `gorm:"index;default:true;"`

	Settings datatypes.JSON `gorm:"not null;default:'{}'::jsonb;"`
}
