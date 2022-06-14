package types

// Package provides version, download url and so on
type Package struct {
	Published int64 `json:"published" bson:"published"`

	Version string `json:"version" bson:"version"`
	URL     string `json:"url" bson:"url"`

	Requirements []string `json:"requirements" bson:"requirements"` // list of needed libraries, etc
}
