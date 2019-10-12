package models

// Pharmacist model struct
type Pharmacist struct {
	Personp personal `json:"person"`

	CompanyName          string  `json:"compagny_name"`
	Turnover             float32 `json:"turnover"`
	EmployeesCount       int     `json:"employees_count"`
	ManagersCount        int     `json:"managers_count"`
	PharmacySurface      float32 `json:"pharmacy_surface"`
	CustomersPerDay      int     `json:"customers_per_day"`
	Workshop             bool    `json:"workshop"`
	TrainingsSold        int     `json:"trainings_sold"`
	WorkshopDaysCount    int     `json:"workshop_days_count"`
	NaturopathyDaysCount int     `json:"naturopathy_days_count"`
}

// Pharmacists many
type Pharmacists []Pharmacist
