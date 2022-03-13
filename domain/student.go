package domain

type Student struct {
	Id         int        `json:"id,omitempty"`
	FirstName  string     `json:"first_name,omitempty"`
	LastName   string     `json:"last_name,omitempty"`
	Identifier string     `json:"identifier,omitempty"`
	modelDate
}
