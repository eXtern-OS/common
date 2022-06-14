package types

import (
	"github.com/eXtern-OS/common/core/interfaces"
)

// ExportedApp provides structure for unified app, this will be our API response
type ExportedApp struct {
	AppType AppType `json:"app_type"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Version string `json:"version"`

	StatsAvailable bool `json:"stats_available"`

	Stats ExportedStats `json:"stats"`

	IconURL     string            `json:"icon_url"`
	HeaderURL   string            `json:"header_url"`
	Screenshots []string          `json:"screenshots"`
	Publisher   ExportedPublisher `json:"publisher"`

	PackageName string  `json:"package_name" bson:"package_name"`
	Package     Package `json:"package" bson:"package"` // only  available for extern apps
}

func ExportApps(income []interfaces.App) []ExportedApp {
	var result []ExportedApp

	for _, x := range income {
		result = append(result, x.Export())
	}

	return result
}
