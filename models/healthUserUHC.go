package models

type HealthUserUHC struct {
	ID           string      `json:"id,omitempty" bson:"id,omitempty"`
	IndividualId string      `json:"individual_id,omitempty" bson:"individual_id,omitempty"`
	MedicalData  MedicalData `json:"medical_data,omitempty" bson:"medical_data,omitempty"`
	Courses      Courses     `json:"courses,omitempty" bson:"courses,omitempty"`
	Donor        Donor       `json:"donor,omitempty" bson:"donor,omitempty"`
}

type Courses struct {
	LocalID string `json:"localId,omitempty" bson:"performed,omitempty"`
}

type MedicalData struct {
	ID string `json:"id,omitempty" bson:"id,omitempty"`
}

type Donor struct {
	ID string `json:"id,omitempty" bson:"id,omitempty"`
}
