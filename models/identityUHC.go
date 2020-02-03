package models

import (
	"time"
)

type IdentityUHC struct {
	ID                string        `json:"id,omitempty" bson:"id,omitempty"`
	IndividualID      string        `json:"individual_id,omitempty" bson:"individual_id,omitempty"`
	Type              string        `json:"type,omitempty" bson:"type,omitempty"`
	Period            string        `json:"period,omitempty" bson:"period,omitempty"`
	Related           string        `json:"related,omitempty" bson:"related,omitempty"`
	Country           string        `json:"country,omitempty" bson:"country,omitempty"`
	OfficialDocType   string        `json:"official_doctype,omitempty" bson:"official_doctype,omitempty"`
	OfficialDocNumber string        `json:"official_doc_number,omitempty" bson:"official_doc_number,omitempty"`
	OfficialName      OfficialName  `json:"official_name,omitempty" bson:"official_name,omitempty"`
	BirthData         Birthdata     `json:"birth_data,omitempty" bson:"birth_data,omitempty"`
	GenderData        Genderdata    `json:"gender_data,omitempty" bson:"gender_data,omitempty"`
	OfficialAddress   []Address     `json:"official_address,omitempty" bson:"official_address,omitempty"`
	RelatedPerson     []interface{} `json:"related_person,omitempty" bson:"related_person,omitempty"`
	Verified          Verified      `json:"verified,omitempty" bson:"verified,omitempty"`
	Temporal          bool          `json:"temporal,omitempty" bson:"temporal,omitempty"`
}

type OfficialName struct {
	ID         string   `json:"id,omitempty" bson:"id,omitempty"`
	GivenName  string   `json:"given_name,omitempty" bson:"given_name,omitempty"`
	MiddleName string   `json:"middle_name,omitempty" bson:"middle_name,omitempty"`
	Surname1   string   `json:"surname1,omitempty" bson:"surname1,omitempty"`
	Surname2   string   `json:"surname2,omitempty" bson:"surname2,omitempty"`
	NameOrder  []string `json:"name_order,omitempty" bson:"name_order,omitempty"`
}

type Birthdata struct {
	Date       time.Time `json:"date,omitempty" bson:"date,omitempty"`
	City       string    `json:"city,omitempty" bson:"city,omitempty"`
	Country    string    `json:"country,omitempty" bson:"country,omitempty"`
	LocationID string    `json:"location_id,omitempty" bson:"location_id,omitempty"`
}

type Genderdata struct {
	Administrative string `json:"administrative,omitempty" bson:"administrative,omitempty"`
	Biological     string `json:"biological,omitempty" bson:"biological,omitempty"`
}

type IdentityUHCResponse struct {
	Code    int
	Message string
	Data    []IdentityUHC
}
