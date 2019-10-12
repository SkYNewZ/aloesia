package models

// Liberal model struct
type Liberal struct {
	Person        personal `json:"person"`
	TrainingsSold int      `json:"trainings_sold"`
}

// Liberals many
type Liberals []Liberal
