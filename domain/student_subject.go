package domain

const (
	StatusInProgress = "in_progress"
	StatusApproved   = "approved"
	StatusReproved   = "reproved"
)

type StudentSubject struct {
	Id        int        `json:"id,omitempty"`
	IdStudent int        `json:"id_student,omitempty"`
	IdSubject int        `json:"id_subject,omitempty"`
	Frequency float64    `json:"frequency,omitempty"`
	Status    string     `json:"status,omitempty"`
	modelDate
}

type StudentSubjectWithSubject struct {
	StudentSubject
	Subject Subject `json:"subject,omitempty"`
}
