package app

// Type is only needed for desktop app store app to understand differences in installing
type Type string

const (
	ExternApp  = "extern"
	SnapApp    = "snap"
	FlatpakApp = "flatpak"
)
