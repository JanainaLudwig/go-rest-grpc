package subject

import (
	"context"
	"grpc-rest/core"
	"time"
)

type Subject struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func FetchAll(ctx context.Context) ([]Subject, error) {
	r := Repository{db: core.DB}

	return r.FetchAll(ctx)
}

func Create(ctx context.Context, std *Subject) (int, error) {
	r := Repository{db: core.DB}

	return r.Insert(ctx, std)
}
