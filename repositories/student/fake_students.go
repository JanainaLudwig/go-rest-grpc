package student

import (
	"github.com/bxcodec/faker/v3"
	"grpc-rest/domain"
	"time"
)

func fakeStudents() []domain.Student {
	var stds []domain.Student
	for i := 0; i < 100; i++ {
		timeNow := time.Now()
		stds = append(stds, domain.Student{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
			ModelDate: domain.ModelDate{
				CreatedAt: &timeNow,
				UpdatedAt: &timeNow,
			},
		})
	}

	return stds
}
