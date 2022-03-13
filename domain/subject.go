package domain

type Subject struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	modelDate
}
