package models

// Institution model struct
type Institution struct {
	Person Personal `json:"person"`

	TrainingsSold int             `json:"trainings_sold"`
	CompanyName   string          `json:"compagny_name"`
	Type          InstitutionType `json:"type"`
	Service       ServiceType     `json:"service"`
}

// Institutions many
type Institutions []Institution
