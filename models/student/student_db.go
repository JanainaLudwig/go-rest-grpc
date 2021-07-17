package student

import (
	"context"
	"database/sql"
	"grpc-rest/core"
	"net/http"
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

		var created, updated sql.NullTime
		err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.Identifier,
			&created,
			&updated,
		)

		if created.Valid {
			student.CreatedAt = &created.Time
		}
		if updated.Valid {
			student.UpdatedAt = &updated.Time
		}

		if err != nil {
			return nil, core.NewError(nil, err, 0)
		}

		students = append(students, student)
	}

	if len(students) == 0 {
		return nil, core.NewError(nil, "Students not found", http.StatusNotFound)
	}

	return students, err
}
