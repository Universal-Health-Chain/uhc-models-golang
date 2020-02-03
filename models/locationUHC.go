package models

type LocationUHC struct {
	ID                     string               `json:"id,omitempty" bson:"id,omitempty"`
	Name                   TextI18n             `json:"name,omitempty" bson:"name,omitempty"`
	Status                 string               `json:"status,omitempty" bson:"status,omitempty"`
	Type                   []string             `json:"type,omitempty" bson:"type,omitempty"`
	PhysicalType           string               `json:"pysical_type,omitempty" bson:"pysical_type,omitempty"`
	Telecom                []Telecom            `json:"telecom,omitempty" bson:"telecom,omitempty"`
	MapWebData             MapWebData           `json:"map_web_data,omitempty" bson:"map_web_data,omitempty"`
	TimeAvailability       string               `json:"time_availability,omitempty" bson:"time_availability,omitempty"`
	Address                Address              `json:"address,omitempty" bson:"address,omitempty"`
	PartOfID               string               `json:"part_of_id,omitempty" bson:"part_of_id,omitempty"`
	ManagingOrganizationID string               `json:"managing_organization_id,omitempty" bson:"managing_organization_id,omitempty"`
	ExternalIdentifier     []ExternalIdentifier `json:"external_identifier,omitempty" bson:"external_identifier,omitempty"`
	Geojson                Geojson              `json:"geojson,omitempty" bson:"geojson,omitempty"`
}

type TextI18n struct {
	Default string            `json:"default,omitempty" bson:"default,omitempty"`
	I18n    map[string]string `json:"i18n,omitempty" bson:"i18n,omitempty"`
}

type MapWebData struct {
	Url string `json:"url,omitempty" bson:"url,omitempty"`
}

type ExternalIdentifier struct {
	System string `json:"system,omitempty" bson:"system,omitempty"`
	Value  string `json:"value,omitempty" bson:"value,omitempty"`
}

type LocationUHCResponse struct {
	Code    int           `json:"code,omitempty" bson:"code,omitempty"`
	Message string        `json:"message,omitempty" bson:"message,omitempty"`
	Data    []LocationUHC `json:"data,omitempty" bson:"data,omitempty"`
	Count   int64         `json:"count,omitempty" bson:"count,omitempty"`
}
