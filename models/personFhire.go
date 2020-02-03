package models

type PersonFhireR4 struct {
	ResourceType string              `bson:"resourceType" json:"resourceType"`
	ID           string              `bson:"id" json:"id"`
	Text         TextFhireR4         `bson:"text" json:"text"`
	Identifier   []IdentifierFhireR4 `bson:"identifier" json:"identifier"`
	Name         []NameFhireR4       `bson:"name" json:"name"`
	Telecom      []TelecomFhireR4    `bson:"telecom" json:"telecom"`
	Gender       string              `bson:"gender" json:"gender"`
	BirthDate    string              `bson:"birthDate" json:"birthDate"`
	Address      []AddressFhireR4    `bson:"address" json:"address"`
	Active       bool                `bson:"active" json:"active"`
	Link         []LinkFhireR4       `bson:"link" json:"link"`
}

type TextFhireR4 struct {
	Status string `bson:"status" json:"status"`
	Div    string `bson:"div" json:"div"`
}
type IdentifierFhireR4 struct {
	Use      string          `bson:"use" json:"use"`
	Type     TypeFhireR4     `bson:"type" json:"type"`
	System   string          `bson:"system" json:"system"`
	Value    string          `bson:"value" json:"value"`
	Period   PeriodFhireR4   `bson:"period" json:"period"`
	Assigner AssignerFhireR4 `bson:"assigner" json:"assigner"`
}

type AssignerFhireR4 struct {
	Display string `bson:"display" json:"display"`
}

type PeriodFhireR4 struct {
	Start string `bson:"start" json:"start"`
}

type TypeFhireR4 struct {
	Coding []struct {
		System string `bson:"system" json:"system"`
		Code   string `bson:"code" json:"code"`
	} `bson:"coding" json:"coding"`
}

type NameFhireR4 struct {
	Use    string   `bson:"use" json:"use"`
	Family string   `bson:"family" json:"family,omitempty"`
	Given  []string `bson:"given" json:"given"`
}

type TelecomFhireR4 struct {
	Use    string `bson:"use" json:"use"`
	System string `bson:"system" json:"system,omitempty"`
	Value  string `bson:"value" json:"value,omitempty"`
}

type AddressFhireR4 struct {
	Use        string   `bson:"use" json:"use"`
	Line       []string `bson:"line" json:"line"`
	City       string   `bson:"city" json:"city"`
	State      string   `bson:"state" json:"state"`
	PostalCode string   `bson:"postalCode" json:"postalCode"`
}

type LinkFhireR4 struct {
	TargetFhireR4 struct {
		Reference string `bson:"reference" json:"reference"`
		Display   string `bson:"display" json:"display"`
	} `bson:"target" json:"target"`
}
