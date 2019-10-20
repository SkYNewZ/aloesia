package models

import "time"

// Personal personal patient informations
type Personal struct {
	FirstName             string    `json:"first_name" firestore:"first_name" mapstructure:"first_name" valid:"required"`
	LastName              string    `json:"last_name" firestore:"last_name" mapstructure:"last_name" valid:"required"`
	Address               string    `json:"address" firestore:"address" mapstructure:"address" valid:"required"`
	PrimaryPhoneNumber    string    `json:"primary_phone_number" mapstructure:"primary_phone_number" firestore:"primary_phone_number" valid:"FRphoneNumber,required"`
	SecondaryPhoneNumber  string    `json:"secondary_phone_number" mapstructure:"secondary_phone_number" firestore:"secondary_phone_number" valid:"FRphoneNumber"`
	PrimaryEmailAddress   string    `json:"primary_email_ddress" mapstructure:"primary_email_ddress" firestore:"primary_email_ddress" valid:"email,required"`
	SecondaryEmailAddress string    `json:"secondary_email_ddress" mapstructure:"secondary_email_ddress" firestore:"secondary_email_ddress" valid:"email"`
	Description           string    `json:"description" mapstructure:"description" firestore:"description"`
	CreatedAt             time.Time `json:"created_at" mapstructure:"created_at" firestore:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" mapstructure:"updated_at" firestore:"updated_at"`
}
