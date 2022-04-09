package subject

import (
	"context"
	"grpc-rest/core"
	"grpc-rest/domain"
)

func FetchAll(ctx context.Context) ([]domain.Subject, error) {
	r := Repository{db: core.DB}

	return r.FetchAll(ctx)
}

func Create(ctx context.Context, std *domain.Subject) (int, error) {
	r := Repository{db: core.DB}

	return r.Insert(ctx, std)
}
