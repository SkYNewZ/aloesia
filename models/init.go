package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/ttacon/libphonenumber"
)

const (
	firestoreUsersCollectionName    = "users"
	firestorePatientsCollectionName = "patients"
	validPhoneNumberCountryCode     = "FR"
)

func init() {
	// Custom validator for FR phonenumber
	govalidator.TagMap["FRphoneNumber"] = govalidator.Validator(func(str string) bool {
		// parse given number to deal with
		num, err := libphonenumber.Parse(str, validPhoneNumberCountryCode)
		if err != nil {
			return false
		}
		// test if given phone number is valid
		return libphonenumber.IsValidNumberForRegion(num, validPhoneNumberCountryCode)
	})

	// Custom validator for valid role
	govalidator.TagMap["role"] = govalidator.Validator(func(str string) bool {
		return str == SuperAdminRole || str == AdminRole
	})
}
