package models

type WorldDataUHC struct {
	ID                 string          `json:"id" bson:"id"`
	ExternalIdentifier []ExternalId    `json:"external_identifier" bson:"external_identifier"`
	Level              string          `json:"level" bson:"level"`
	LevelType          string          `json:"level_type" bson:"level_type"`
	Sub                []ExternalIdSub `json:"sub" bson:"sub"`
	Root               []ExternalId    `json:"root" bson:"root"`
	Flag               Flag            `json:"flag" bson:"flag"`
	Locale             Locale          `json:"locale" bson:"locale"`
	Geojson            Geojson         `json:"geojson" bson:"geojson"`
	Display            Display         `json:"display" bson:"display"`
}

type ExternalId struct {
	System string `json:"system,omitempty" bson:"system,omitempty"`
	Value  string `json:"value,omitempty" bson:"value,omitempty"`
}

type ExternalIdSub struct {
	System string   `json:"system,omitempty" bson:"system,omitempty"`
	Value  []string `json:"value,omitempty" bson:"value,omitempty"`
}

type Flag struct {
	URL string `json:"url,omitempty" bson:"url,omitempty"`
}

type Locale struct {
	CallingCode       []string `json:"calling_code,omitempty" bson:"calling_code,omitempty"`
	Currency          []string `json:"currency,omitempty" bson:"currency,omitempty"`
	OfficialLanguages []string `json:"official_languages,omitempty" bson:"official_languages,omitempty"`
}

type Display struct {
	Localname  string                `json:"localname,omitempty" bson:"localname,omitempty"`
	I18NRegion map[string]I18NRegion `json:"i18n,omitempty" bson:"i18n,omitempty"`
}

type I18NRegion struct {
	Common string `json:"common,omitempty" bson:"common,omitempty"`
}

type WorldDataUHCResponse struct {
	Code    int            `json:"code,omitempty" bson:"code,omitempty"`
	Message string         `json:"message,omitempty" bson:"message,omitempty"`
	Data    []WorldDataUHC `json:"data,omitempty" bson:"data,omitempty"`
}
