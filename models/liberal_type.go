package models

// LiberalType type of Liberal
type LiberalType struct {
	ID       string   `json:"id"`
	Label    string   `json:"label"`
	Liberals Liberals `json:"liberals"`
}
