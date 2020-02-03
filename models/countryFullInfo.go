package models

type CountryFullInfo struct {
	Country    Display                  `json:"country_display,omitempty" bson:"country_display,omitempty"`
	CountryIso string                   `json:"county_iso,omitempty" bson:"county_iso,omitempty"`
	States     map[string]StateFullInfo `json:"states,omitempty" bson:"states,omitempty"`
}

type StateFullInfo struct {
	StateDisplay Display                     `json:"state_display,omitempty" bson:"state_display,omitempty"`
	StateIso     string                      `json:"state_iso,omitempty" bson:"state_iso,omitempty"`
	Districts    map[string]DistrictFullInfo `json:"districts,omitempty" bson:"districts,omitempty"`
}

type DistrictFullInfo struct {
	DistrictDisplay Display `json:"district_display,omitempty" bson:"district_display,omitempty"`
	DistrictIso     string  `json:"district_iso,omitempty" bson:"district_iso,omitempty"`
}

type CountryFullInfoResponse struct {
	Code    int             `json:"code,omitempty" bson:"code,omitempty"`
	Message string          `json:"message,omitempty" bson:"message,omitempty"`
	Data    CountryFullInfo `json:"data,omitempty" bson:"data,omitempty"`
}
