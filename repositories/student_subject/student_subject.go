package student_subject

import (
	"context"
	"grpc-rest/core"
	"grpc-rest/domain"
)

func FetchAll(ctx context.Context) ([]domain.StudentSubject, error) {
	r := Repository{db: core.DB}

	return r.FetchAll(ctx)
}

func FetchByStudentSubjectId(ctx context.Context, idStudent int) ([]domain.StudentSubjectWithSubject, error) {
	r := Repository{db: core.DB}

	return r.FetchByStudentSubjectId(ctx, idStudent)
}

func Create(ctx context.Context, std *domain.StudentSubject) (int, error) {
	r := Repository{db: core.DB}

	return r.Insert(ctx, std)
}
