package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/ttacon/libphonenumber"
)

const (
	firestoreUsersCollectionName        = "users"
	firestoreLiberalTypesCollectionName = "liberal_types"
	firestorePatientsCollectionName     = "patients"
	validPhoneNumberCountryCode         = "FR"
)

func init() {
	// Add your own struct validation tags
	govalidator.TagMap["FRphoneNumber"] = govalidator.Validator(func(str string) bool {
		// parse given number to deal with
		num, err := libphonenumber.Parse(str, validPhoneNumberCountryCode)
		if err != nil {
			return false
		}
		// test if given phone number is valid
		return libphonenumber.IsValidNumberForRegion(num, validPhoneNumberCountryCode)
	})
}
