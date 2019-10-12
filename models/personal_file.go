package models

type personal struct {
	ID                    string `json:"id"`
	LastName              string `json:"last_name"`
	FirstName             string `json:"first_name"`
	Address               string `json:"address"`
	PrimaryPhoneNumber    string `json:"primary_phone_number"`
	SecondaryPhoneNumber  string `json:"secondary_phone_number"`
	PrimaryEmailAddress   string `json:"primary_email_ddress"`
	SecondaryEmailAddress string `json:"secondary_email_ddress"`
	Description           string `json:"description"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}
