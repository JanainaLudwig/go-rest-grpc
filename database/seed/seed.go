package seed

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"grpc-rest/models/student"
	"grpc-rest/models/student_subject"
	"grpc-rest/models/subject"
	"sync"
)

func RunSeed(ctx context.Context) {
	subjectsIds := subjects(ctx)
	students(ctx, subjectsIds)
}

func students(ctx context.Context, subjectsIds []int) {
	var stds []student.Student
	for i := 0; i < 3000; i++ {
		stds = append(stds, student.Student{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
		})
	}

	wg := sync.WaitGroup{}
	wg.Add(len(stds))
	for _, std := range stds {
		go func(std student.Student) {
			defer wg.Done()
			id, err := student.Create(ctx, &std)
			if err != nil {
				return
			}

			freq, _ := faker.RandomInt(0, 100)

			for _, idSubject := range subjectsIds {
				student_subject.Create(ctx, &student_subject.StudentSubject{
					IdStudent: id,
					IdSubject: idSubject,
					Frequency: float64(freq[0]),
					Status:    student_subject.StatusInProgress,
				})
			}
		}(std)
	}
	wg.Wait()
}

func subjects(ctx context.Context) []int {
	stds := []subject.Subject{}
	for i := 0; i < 100; i++ {
		stds = append(stds, subject.Subject{
			Name: faker.Word(),
		})
	}

	var ids []int

	for _, std := range stds {
		create, err := subject.Create(ctx, &std)
		if err != nil {
			return nil
		}

		ids = append(ids, create)
	}

	return ids
}
