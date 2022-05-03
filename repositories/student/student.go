package student

import (
	"context"
	"github.com/google/uuid"
	"grpc-rest/core"
	"grpc-rest/domain"
)

func FetchAll(ctx context.Context) ([]domain.Student, error) {
	//r := Repository{db: core.DB}
	//
	//return r.FetchAll(ctx)

	return fakeStudents(), nil
}

func FetchById(ctx context.Context, id int) (*domain.Student, error) {
	r := Repository{db: core.DB}

	return r.FetchById(ctx, id)
}

func Create(ctx context.Context, std *domain.Student) (int, error) {
	r := Repository{db: core.DB}

	std.Identifier = uuid.New().String()

	return r.Insert(ctx, std)
}

func Update(ctx context.Context, std *domain.Student) error {
	r := Repository{db: core.DB}

	return r.Update(ctx, std)
}

func Delete(ctx context.Context, id int) error {
	r := Repository{db: core.DB}

	return r.Delete(ctx, id)
}
