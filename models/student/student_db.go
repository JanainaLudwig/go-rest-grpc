package student

import (
	"context"
	"database/sql"
	"grpc-rest/core"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) FetchAll(ctx context.Context) ([]Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, first_name, last_name, identifier, created_at, updated_at FROM students")
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	var students []Student
	for rows.Next() {
		var student Student

		err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.Identifier,
			&student.CreatedAt,
			&student.UpdatedAt,
		)

		if err != nil {
			return nil, core.NewError(nil, err, 0)
		}
	}

	return students, err
}

