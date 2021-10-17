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

func scanStudent(rows *sql.Rows) (*Student, error) {
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
	if err != nil {
		return nil, err
	}

	if created.Valid {
		student.CreatedAt = &created.Time
	}
	if updated.Valid {
		student.UpdatedAt = &updated.Time
	}

	return &student, nil
}

func (r *Repository) FetchAll(ctx context.Context) ([]Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, first_name, last_name, identifier, created_at, updated_at FROM students")
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	defer core.DbClose(rows)

	var students []Student
	for rows.Next() {
		student, err := scanStudent(rows)

		if err != nil {
			return nil, core.NewError(nil, err, 0)
		}

		students = append(students, *student)
	}

	if len(students) == 0 {
		return nil, core.NewError(nil, "Students not found", http.StatusNotFound)
	}

	return students, err
}

func (r *Repository) Insert(ctx context.Context, std *Student) (int, error) {
	res, err := r.db.ExecContext(ctx, "INSERT INTO students (first_name, last_name, identifier) VALUES (?, ?, ?)",
		std.FirstName, std.LastName, std.Identifier)
	if err != nil {
		return 0, core.NewError(nil, err, 0)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, core.WrapError(err)
	}

	return int(id), err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM students WHERE id=?", id)
	if err != nil {
		return core.NewError(nil, err, 0)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return core.WrapError(err)
	}

	if affected == 0 {
		return core.NewError(id, "Student not found", http.StatusNotFound)
	}

	return nil
}

func (r *Repository) FetchById(ctx context.Context, id int) (*Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, first_name, last_name, identifier, created_at, updated_at FROM students")
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	defer core.DbClose(rows)

	if !rows.Next() {
		return nil, core.NewError(nil, "Students not found", http.StatusNotFound)
	}

	student, err := scanStudent(rows)

	return student, err
}
