package student

import (
	"context"
	"grpc-rest/core"
	"time"
	"github.com/google/uuid"
)

type Student struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Identifier string    `json:"identifier"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func FetchAll(ctx context.Context) ([]Student, error) {
	r := Repository{db: core.DB}

	return r.FetchAll(ctx)
}

func Insert(ctx context.Context, std *Student) (int, error) {
	r := Repository{db: core.DB}

	std.Identifier = uuid.New().String()

	return r.Insert(ctx, std)
}

func Delete(ctx context.Context, id int) error {
	r := Repository{db: core.DB}

	return r.Delete(ctx, id)
}
