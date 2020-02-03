package models

type ProcedureUHC struct {
	ID              string `json:"id,omitempty" bson:"id,omitempty"`
	Performed       string `json:"performed,omitempty" bson:"performed,omitempty"`
	Performer       string `json:"performer,omitempty" bson:"performer,omitempty"`
	State           string `json:"state,omitempty" bson:"state,omitempty"`
	Success         string `json:"success,omitempty" bson:"success,omitempty"`
	Interruption    string `json:"interruption,omitempty" bson:"interruption,omitempty"`
	Complication    string `json:"complication,omitempty" bson:"complication,omitempty"`
	Adversereaction string `json:"adversereaction,omitempty" bson:"adversereaction,omitempty"`
	Note            string `json:"note,omitempty" bson:"note,omitempty"`
	Outcome         string `json:"outcome,omitempty" bson:"outcome,omitempty"`
}

type Phlebotomy struct {
	Arm    string `json:"arm,omitempty" bson:"arm,omitempty"`
	Issue  string `json:"issue,omitempty" bson:"issue,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
}
