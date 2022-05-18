package student_subject

import (
	"context"
	"grpc-rest/core"
	"grpc-rest/domain"
	"time"
)

func FetchAll(ctx context.Context) ([]domain.StudentSubject, error) {
	r := Repository{db: core.DB}

	return r.FetchAll(ctx)
}

func FetchByStudentSubjectId(ctx context.Context, idStudent int) ([]domain.StudentSubjectWithSubject, error) {
	//r := Repository{db: core.DB}
	//
	//return r.FetchByStudentSubjectId(ctx, idStudent)
	timeNow := time.Now()
	sampleSubject := domain.StudentSubjectWithSubject{
		StudentSubject: domain.StudentSubject{
			Id:        1,
			IdStudent: 1,
			IdSubject: 1,
			Frequency: 95,
			Status:    domain.StatusInProgress,
			ModelDate: domain.ModelDate{
				CreatedAt: &timeNow,
				UpdatedAt: &timeNow,
			},
		},
		Name:           "Math",
	}

	if idStudent == 1 {
		return []domain.StudentSubjectWithSubject{sampleSubject}, nil
	}

	var subjects []domain.StudentSubjectWithSubject
	for i := 0; i < 50; i++ {
		subjects = append(subjects, sampleSubject)
	}

	return subjects, nil
}

func Create(ctx context.Context, std *domain.StudentSubject) (int, error) {
	r := Repository{db: core.DB}

	return r.Insert(ctx, std)
}
