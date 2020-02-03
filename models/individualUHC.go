package models

import (
	"time"
)

type IndividualUHC struct {
	Type         int           `json:"type,omitempty" bson:"type,omitempty"`
	Display      string        `json:"display,omitempty" bson:"display,omitempty"`
	ID           string        `json:"id,omitempty" bson:"id,omitempty"`
	PublicId     []PublicId    `json:"public_id,omitempty" bson:"public_id,omitempty"`
	IdentityIDs  []string      `json:"identity_ids,omitempty" bson:"identity_ids,omitempty"`
	Verification Verification  `json:"verification,omitempty" bson:"verification,omitempty"`
	PersonalData PersonalData  `json:"personal_data,omitempty" bson:"personal_data,omitempty"`
	EncounterIds []string      `json:"encounter_ids,omitempty" bson:"encounter_ids,omitempty"`
	Profiles     []interface{} `json:"profiles,omitempty" bson:"profiles,omitempty"`
}

type PublicId struct {
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Value string `json:"value,omitempty" bson:"vale,omitempty"`
}

type Verified struct {
	Done     bool     `json:"done,omitempty" bson:"done,omitempty"`
	Result   bool     `json:"result,omitempty" bson:"result,omitempty"`
	Attested Attested `json:"attested,omitempty" bson:"attested,omitempty"`
	At       At       `json:"at,omitempty" bson:"at,omitempty"`
}

type Attested struct {
	ProfessionalID string `json:"professional_id,omitempty" bson:"professional_id,omitempty"`
	PersonID       string `json:"person_id,omitempty" bson:"person_id,omitempty"`
	OrgID          string `json:"org_id,omitempty" bson:"org_id,omitempty"`
}

type At struct {
	Place     Place     `json:"place,omitempty" bson:"place,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type Place struct {
	LocationID string `json:"location_id,omitempty" bson:"location_id,omitempty"`
	Display    string `json:"display,omitempty" bson:"display,omitempty"`
}

type Verification struct {
	Type     string   `json:"type,omitempty" bson:"type,omitempty"`
	Done     bool     `json:"done,omitempty" bson:"done,omitempty"`
	Result   string   `json:"result,omitempty" bson:"result,omitempty"`
	Attested Attested `json:"attested,omitempty" bson:"attested,omitempty"`
	At       At       `json:"at,omitempty" bson:"at,omitempty"`
}

type PersonalData struct {
	MainIdentityId string     `json:"main_identity_id,omitempty" bson:"main_identity_id,omitempty"`
	Address        []Address  `json:"address,omitempty" bson:"address,omitempty"`
	CommonName     CommonName `json:"common_name,omitempty" bson:"common_name,omitempty"`
	Telecom        []Telecom  `json:"telecom,omitempty" bson:"telecom,omitempty"`
	Photo          string     `json:"photo,omitempty" bson:"photo,omitempty"`
	Ethic          string     `json:"ethic,omitempty" bson:"ethic,omitempty"`
	Religion       string     `json:"religion,omitempty" bson:"religion,omitempty"`
}

type Address struct {
	Type        string        `json:"type,omitempty" bson:"type,omitempty"`
	Use         string        `json:"use,omitempty" bson:"use,omitempty"`
	Text        string        `json:"text,omitempty" bson:"text,omitempty"`
	Comment     string        `json:"comment,omitempty" bson:"comment,omitempty"`
	Line        []interface{} `json:"line,omitempty" bson:"line,omitempty"`
	Census      string        `json:"census,omitempty" bson:"census,omitempty"`
	PostalCode  string        `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
	City        string        `json:"city,omitempty" bson:"city,omitempty"`
	DistrictIso string        `json:"district_iso,omitempty" bson:"district_iso,omitempty"`
	StateIso    string        `json:"state_iso,omitempty" bson:"state_iso,omitempty"`
	CountryIso  string        `json:"country_iso,omitempty" bson:"country_iso,omitempty"`
	Period      Period        `json:"period,omitempty" bson:"period,omitempty"`
}

type Period struct {
	Start time.Time `json:"start,omitempty" bson:"start,omitempty"`
	End   time.Time `json:"end,omitempty" bson:"end,omitempty"`
}
type CommonName struct {
	NickName string `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
}

type Organization struct {
	ID        string `json:"id,omitempty" bson:"id,omitempty"`
	LocalCard string `json:"local_card,omitempty" bson:"local_card,omitempty"`
}

type Telecom struct {
	Use    string `bson:"use" json:"use"`
	System string `bson:"system" json:"system,omitempty"`
	Value  string `bson:"value" json:"value,omitempty"`
}

type IndividualUHCResponse struct {
	Code    int             `json:"code,omitempty" bson:"code,omitempty"`
	Message string          `json:"message,omitempty" bson:"message,omitempty"`
	Data    []IndividualUHC `json:"data,omitempty" bson:"data,omitempty"`
	Count   int64           `json:"count,omitempty" bson:"count,omitempty"`
}
