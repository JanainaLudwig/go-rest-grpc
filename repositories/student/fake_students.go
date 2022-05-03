package student

import (
	"github.com/bxcodec/faker/v3"
	"grpc-rest/domain"
)

func fakeStudents() []domain.Student {
	var stds []domain.Student
	for i := 0; i < 3000; i++ {
		stds = append(stds, domain.Student{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
		})
	}

	return stds
}
