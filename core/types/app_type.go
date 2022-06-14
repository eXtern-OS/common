package types

// AppType is only needed for desktop app store app to understand differences in installing
type AppType string

const (
	ExternApp  AppType = "extern"
	SnapApp    AppType = "snap"
	FlatpakApp AppType = "flatpak"
)
