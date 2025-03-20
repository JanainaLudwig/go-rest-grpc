package student

import (
	"context"
	"grpc-rest/core"
	"grpc-rest/domain"
	"time"

	"github.com/google/uuid"
)

func FetchAll(ctx context.Context) ([]domain.Student, error) {
	//r := Repository{db: core.DB}
	//
	//return r.FetchAll(ctx)

	return fakeStudents(), nil
}


func FetchAll2(ctx context.Context) ([]domain.Student, error) {
	return fakeStudents(), nil
}

func FetchById(ctx context.Context, id int) (*domain.Student, error) {
	//r := Repository{db: core.DB}

	//return r.FetchById(ctx, id)
	timeNow := time.Now()
	return &domain.Student{
		Id:         1,
		FirstName:  "John",
		LastName:   "Doe",
		Identifier: "34546468",
		ModelDate:  domain.ModelDate{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}, nil
}

func FetchById2(ctx context.Context, id int) (*domain.Student, error) {
	//r := Repository{db: core.DB}

	//return r.FetchById(ctx, id)
	timeNow := time.Now()
	return &domain.Student{
		Id:         1,
		FirstName:  "John",
		LastName:   "Doe",
		Identifier: "34546468",
		ModelDate:  domain.ModelDate{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}, nil
}
func FetchById3(ctx context.Context, id int) (*domain.Student, error) {
	//r := Repository{db: core.DB}

	//return r.FetchById(ctx, id)
	timeNow := time.Now()
	return &domain.Student{
		Id:         1,
		FirstName:  "John",
		LastName:   "Doe",
		Identifier: "34546468",
		ModelDate:  domain.ModelDate{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}, nil
}

func FetchById4(ctx context.Context, id int) (*domain.Student, error) {
	//r := Repository{db: core.DB}

	//return r.FetchById(ctx, id)
	timeNow := time.Now()
	return &domain.Student{
		Id:         1,
		FirstName:  "John",
		LastName:   "Doe",
		Identifier: "34546468",
		ModelDate:  domain.ModelDate{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}, nil
}

func FetchById5(ctx context.Context, id int) (*domain.Student, error) {
	//r := Repository{db: core.DB}

	//return r.FetchById(ctx, id)
	timeNow := time.Now()
	return &domain.Student{
		Id:         1,
		FirstName:  "John",
		LastName:   "Doe",
		Identifier: "34546468",
		ModelDate:  domain.ModelDate{
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}, nil
}
func Create(ctx context.Context, std *domain.Student) (int, error) {
	//r := Repository{db: core.DB}

	std.Identifier = uuid.New().String()

	return 1, nil
	//return r.Insert(ctx, std)
}

func Update(ctx context.Context, std *domain.Student) error {
	r := Repository{db: core.DB}

	return r.Update(ctx, std)
}

func Delete(ctx context.Context, id int) error {
	r := Repository{db: core.DB}

	return r.Delete(ctx, id)
}
