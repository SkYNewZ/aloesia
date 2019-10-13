package models

// Liberal model struct
type Liberal struct {
	Person        Personal `json:"person"`
	TrainingsSold int      `json:"trainings_sold"`
}

// Liberals many
type Liberals []Liberal
