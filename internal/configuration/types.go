package configuration

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Configuration struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	SiteID  uuid.UUID `json:"site_id"`
	Name    string    `json:"name"`
	Plugins []Plugin  `json:"plugins"`
}

type Plugin struct {
	ID        uuid.UUID      `json:"id"`
	Namespace string         `json:"namespace"`
	Name      string         `json:"name"`
	Settings  datatypes.JSON `json:"settings"`
}
