package models

type PatientFhireR4 struct {
	ResourceType         string                      `bson:"resourceType" json:"resourceType"`
	ID                   string                      `bson:"id" json:"id"`
	Text                 TextFhireR4                 `bson:"text" json:"text"`
	Identifier           []IdentifierFhireR4         `bson:"identifier" json:"identifier"`
	Active               bool                        `bson:"active" json:"active"`
	Name                 []NameFhireR4               `bson:"name" json:"name"`
	Gender               string                      `bson:"gender" json:"gender"`
	Gender_              GenderFhireR4_              `bson:"_gender" json:"_gender"`
	Photo                []PhotoFhireR4              `bson:"photo" json:"photo"`
	ManagingOrganization ManagingOrganizationFhireR4 `bson:"managingOrganization" json:"managingOrganization"`
	Link                 []LinkFhireR4               `bson:"link" json:"link"`
}

type PhotoFhireR4 struct {
	ContentType string `bson:"contentType" json:"contentType"`
	Data        string `bson:"data" json:"data"`
}

type ManagingOrganizationFhireR4 struct {
	Reference string `bson:"reference" json:"reference"`
	Display   string `bson:"display" json:"display"`
}

type GenderFhireR4_ struct {
	Extension []ExtensionFhireR4 `bson:"extension" json:"extension"`
}

type ExtensionFhireR4 struct {
	URL                  string                      `bson:"url" json:"url"`
	ValueCodeableConcept ValueCodeableConceptFhireR4 `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
}

type ValueCodeableConceptFhireR4 struct {
	Coding []struct {
		System  string `bson:"system" json:"system"`
		Code    string `bson:"code" json:"code"`
		Display string `bson:"display" json:"display"`
	} `json:"coding"`
}
