package models

type IndividualUHCWithIdentity struct {
	Individual IndividualUHC `json:"individual,omitempty" bson:"individual,omitempty"`
	Identity   []IdentityUHC `json:"identity,omitempty" bson:"identity,omitempty"`
}

type IndividualUHCWithIdentityResponse struct {
	Code    int                         `json:"code,omitempty" bson:"code,omitempty"`
	Message string                      `json:"message,omitempty" bson:"message,omitempty"`
	Data    []IndividualUHCWithIdentity `json:"data,omitempty" bson:"data,omitempty"`
	Count   int64                       `json:"count,omitempty" bson:"count,omitempty"`
}
