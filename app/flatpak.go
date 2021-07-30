package app

import (
	"github.com/eXtern-OS/common/publisher"
)

/* Sample {
	"flatpakAppId":"com.play0ad.zeroad",
	"name":"0 A.D.",
	"summary": "Real-Time Strategy Game of Ancient Warfare",
	"currentReleaseVersion":"0.0.24",
	"currentReleaseDate":"2021-07-21",
	"iconDesktopUrl":"https://dl.flathub.org/repo/appstream/x86_64/icons/128x128/com.play0ad.zeroad.png",
	"iconMobileUrl":"https://dl.flathub.org/repo/appstream/x86_64/icons/128x128/com.play0ad.zeroad.png",
	"inStoreSinceDate":"2017-04-18T04:14:01Z"
}
*/

// Flatpak provides structure for raw response from API
type Flatpak struct {
	FlatpakAppId          string `json:"flatpakAppId" bson:"flatpakAppId"`
	Name                  string `json:"name" bson:"name"`
	Summary               string `json:"summary" bson:"summary"`
	CurrentReleaseVersion string `json:"currentReleaseVersion" bson:"currentReleaseVersion"`
	CurrentReleaseDate    string `json:"currentReleaseDate" bson:"currentReleaseDate"`
	IconDesktopUrl        string `json:"iconDesktopUrl" bson:"iconDesktopUrl"`
	IconMobileUrl         string `json:"iconMobileUrl" bson:"iconMobileUrl"`
	InStoreSinceDate      string `json:"inStoreSinceDate" bson:"inStoreSinceDate"`
}

func (f *Flatpak) Export() ExportedApp {
	return ExportedApp{
		AppType:        FlatpakApp,
		Name:           f.Name,
		Description:    f.Summary,
		Version:        f.CurrentReleaseVersion,
		StatsAvailable: false,
		IconURL:        f.IconDesktopUrl,
		Publisher:      publisher.NewPublisher(f.Name),
		PackageName:    f.FlatpakAppId,
	}
}

func (f *Flatpak) IsPaid() bool {
	return false
}
