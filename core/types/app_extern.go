package types

// Extern is an app which was developed for extern OS
type Extern struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	Icon        string   `json:"icon" bson:"icon"`
	Header      string   `json:"header" bson:"header"`
	Screenshots []string `json:"screenshots" bson:"screenshots"`

	Publisher Publisher `json:"publisher" bson:"publisher"`

	Latest   Package `json:"latest" bson:"latest"`
	Packages Package `json:"packages" bson:"packages"`

	Payment Payment `json:"payment" bson:"payment"`

	StatsAvailable bool `json:"stats_available" bson:"stats_available"`

	Stats Stats `json:"stats" bson:"stats"`
}

func (e *Extern) Export() ExportedApp {
	return ExportedApp{
		AppType:        ExternApp,
		Name:           e.Name,
		Description:    e.Description,
		Version:        e.Latest.Version,
		StatsAvailable: e.StatsAvailable,
		Stats:          e.Stats.Export(),
		IconURL:        e.Icon,
		HeaderURL:      e.Header,
		Publisher:      e.Publisher.Export(),
		Package:        e.Latest,
		Screenshots:    e.Screenshots,
	}
}

func (e *Extern) IsPaid() bool {
	return e.Payment != Free
}
