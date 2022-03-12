package student_subject

import (
	"context"
	"grpc-rest/core"
	"time"
)

const (
	StatusInProgress = "in_progress"
	StatusApproved = "approved"
	StatusReproved = "reproved"
)

type StudentSubject struct {
	Id        int        `json:"id"`
	IdStudent int        `json:"id_student"`
	IdSubject int        `json:"id_subject"`
	Frequency float64    `json:"frequency"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func FetchAll(ctx context.Context) ([]StudentSubject, error) {
	r := Repository{db: core.DB}

	return r.FetchAll(ctx)
}

func Create(ctx context.Context, std *StudentSubject) (int, error) {
	r := Repository{db: core.DB}

	return r.Insert(ctx, std)
}
