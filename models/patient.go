package models

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/ttacon/libphonenumber"

	"github.com/SkYNewZ/aloesia/config"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

// Patient model struct
type Patient struct {
	ID                   string   `json:"id" firestore:"id" mapstructure:"id" valid:"uuidv4"`
	PersonalData         Personal `json:"personal_data" mapstructure:"personal_data" firestore:"personal_data"`
	BirthdayDate         string   `json:"birthday_date" mapstructure:"birthday_date" firestore:"birthday_date"`
	FamilyStatus         string   `json:"family_status" mapstructure:"family_status" firestore:"family_status"`
	ProfessionalActivity string   `json:"professional_activity" mapstructure:"professional_activity" firestore:"professional_activity"`
	Hobbies              string   `json:"hobbies" mapstructure:"hobbies" firestore:"hobbies"`
	VisitReason          string   `json:"visit_reason" mapstructure:"visit_reason" firestore:"visit_reason"`
	Vaccinated           bool     `json:"vaccinated" mapstructure:"vaccinated" firestore:"vaccinated"`
	DentalAmalgam        bool     `json:"dental_amalgam" mapstructure:"dental_amalgam" firestore:"dental_amalgam"`
	Smoking              bool     `json:"smoking" mapstructure:"smoking" firestore:"smoking"`
	ToxicEnvironment     bool     `json:"toxic_environment" mapstructure:"toxic_environment" firestore:"toxic_environment"`
	SurgeryOperation     bool     `json:"surgery_operation" mapstructure:"surgery_operation" firestore:"surgery_operation"`
	Anesthesia           bool     `json:"anesthesia" mapstructure:"anesthesia" firestore:"anesthesia"`
	Surgery              string   `json:"surgery" mapstructure:"surgery" firestore:"surgery"`
	MedicalMonitoring    string   `json:"medical_monitoring" mapstructure:"medical_monitoring" firestore:"medical_monitoring"`
	FamilyBackground     string   `json:"family_background" mapstructure:"family_background" firestore:"family_background"`
}

// Patients many
type Patients []Patient

// GetAllPatients return all patients from database
func GetAllPatients() (*Patients, error) {
	iter := config.Firestore().Collection(firestorePatientsCollectionName).Documents(context.Background())
	data := make(Patients, 0)
	// if there is no data neither collection
	if iter == nil {
		return &data, nil
	}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var patient Patient
		err = mapstructure.Decode(doc.Data(), &patient)
		if err != nil {
			return nil, err
		}
		data = append(data, patient)
	}
	return &data, nil
}

// GetOnePatient return one patient by given id
func GetOnePatient(id string) (*Patient, error) {
	dsnap, err := config.Firestore().Collection(firestorePatientsCollectionName).Doc(id).Get(context.Background())
	if err != nil {
		return &Patient{}, err
	}
	var patient Patient
	err = mapstructure.Decode(dsnap.Data(), &patient)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

// CreatePatient create new patient in datatbase
func CreatePatient(patient *Patient) error {
	patient.ID = uuid.NewV4().String()

	// format phone number
	// required primary phone number
	num, _ := libphonenumber.Parse(patient.PersonalData.PrimaryPhoneNumber, validPhoneNumberCountryCode)
	patient.PersonalData.PrimaryPhoneNumber = libphonenumber.Format(num, libphonenumber.E164)

	// optional secondary phone number
	num, _ = libphonenumber.Parse(patient.PersonalData.SecondaryPhoneNumber, validPhoneNumberCountryCode)
	patient.PersonalData.SecondaryPhoneNumber = libphonenumber.Format(num, libphonenumber.E164)

	// date
	patient.PersonalData.CreatedAt = time.Now()
	patient.PersonalData.UpdatedAt = time.Now()

	_, err := config.Firestore().Collection(firestorePatientsCollectionName).Doc(patient.ID).Set(context.Background(), &patient)
	return err
}

// DeletePatient delete given patient form database
func DeletePatient(id string) error {
	_, err := config.Firestore().Collection(firestorePatientsCollectionName).Doc(id).Delete(context.Background())
	return err
}
