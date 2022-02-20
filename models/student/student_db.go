package student

import (
	"context"
	"database/sql"
	"grpc-rest/core"
)

type Repository struct {
	db *sql.DB
}

const (
	ErrorStudentNotFound = "Student not found"
)

func scanDefaultStudent(rows *sql.Rows) (*Student, error) {
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
		student, err := scanDefaultStudent(rows)

		if err != nil {
			return nil, core.NewError(nil, err, 0)
		}

		students = append(students, *student)
	}

	if len(students) == 0 {
		return nil, core.NotFoundError(nil, "Students not found")
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

func (r *Repository) Update(ctx context.Context, std *Student) error {
	res, err := r.db.ExecContext(ctx, "UPDATE students SET first_name=?, last_name=? WHERE id=?",
		std.FirstName, std.LastName, std.Id)
	if err != nil {
		return core.NewError(nil, err, 0)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return core.WrapError(err)
	}

	if rows == 0 {
		return core.NotFoundError(std.Id, ErrorStudentNotFound)
	}

	return nil
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
		return core.NotFoundError(id, ErrorStudentNotFound)
	}

	return nil
}

func (r *Repository) FetchById(ctx context.Context, id int) (*Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, first_name, last_name, identifier, created_at, updated_at FROM students WHERE id=?", id)
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	defer core.DbClose(rows)

	if !rows.Next() {
		return nil, core.NotFoundError(id, ErrorStudentNotFound)
	}

	student, err := scanDefaultStudent(rows)

	return student, err
}
