package models

type EncounterUHC struct {
	ID          string `json:"id,omitempty" bson:"id,omitempty"`
	Status      string `json:"status,omitempty" bson:"status,omitempty"`
	Appointment string `json:"appointment,omitempty" bson:"appointment,omitempty"`
	Period      string `json:"period,omitempty" bson:"period,omitempty"`
	Length      string `json:"length,omitempty" bson:"length,omitempty"`
	Location    string `json:"location,omitempty" bson:"location,omitempty"`
}

type Appointment struct {
	ID string `json:"id,omitempty" bson:"id,omitempty"`
}
