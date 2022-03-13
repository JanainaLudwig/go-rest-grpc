package seed

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"grpc-rest/domain"
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
	var stds []domain.Student
	for i := 0; i < 3000; i++ {
		stds = append(stds, domain.Student{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
		})
	}

	wg := sync.WaitGroup{}
	wg.Add(len(stds))
	for _, std := range stds {
		go func(std domain.Student) {
			defer wg.Done()
			id, err := student.Create(ctx, &std)
			if err != nil {
				return
			}

			for _, idSubject := range subjectsIds {
				freq, _ := faker.RandomInt(0, 100)

				student_subject.Create(ctx, &domain.StudentSubject{
					IdStudent: id,
					IdSubject: idSubject,
					Frequency: float64(freq[0]),
					Status:    getRandomStatus(),
				})
			}
		}(std)
	}
	wg.Wait()
}

func getRandomStatus() string {
	statusList := []string{domain.StatusInProgress, domain.StatusReproved, domain.StatusApproved}
	statusIndex, _ := faker.RandomInt(0, len(statusList) - 1)

	return statusList[statusIndex[0]]
}

func subjects(ctx context.Context) []int {
	stds := []domain.Subject{}
	for i := 0; i < 100; i++ {
		stds = append(stds, domain.Subject{
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
