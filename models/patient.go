package models

// Patient model struct
type Patient struct {
	Person Personal `json:"person"`

	BirthdayDate         string `json:"birthday_date"`
	FamilyStatus         string `json:"family_status"`
	ProfessionalActivity string `json:"professional_activity"`
	Hobbies              string `json:"hobbies"`
	VisitReason          string `json:"visit_reason"`
	Vaccinated           bool   `json:"vaccinated"`
	DentalAmalgam        bool   `json:"dental_amalgam"`
	Smoking              bool   `json:"smoking"`
	ToxicEnvironment     bool   `json:"toxic_environment"`
	SurgeryOperation     bool   `json:"surgery_operation"`
	Anesthesia           bool   `json:"anesthesia"`
	Surgery              string `json:"surgery"`
	MedicalMonitoring    string `json:"medical_monitoring"`
	FamilyBackground     string `json:"family_background"`
}

// Patients many
type Patients []Patient
