package student_subject

import (
	"context"
	"database/sql"
	"grpc-rest/core"
	"grpc-rest/domain"
)

type Repository struct {
	db *sql.DB
}

func scanDefaultStudentSubject(rows *sql.Rows) (*domain.StudentSubject, error) {
	var subject domain.StudentSubject

	var created, updated sql.NullTime
	err := rows.Scan(
		&subject.Id,
		&subject.IdStudent,
		&subject.IdSubject,
		&subject.Frequency,
		&subject.Status,
		&created,
		&updated,
	)
	if err != nil {
		return nil, err
	}

	if created.Valid {
		subject.CreatedAt = &created.Time
	}
	if updated.Valid {
		subject.UpdatedAt = &updated.Time
	}

	return &subject, nil
}

func (r *Repository) FetchAll(ctx context.Context) ([]domain.StudentSubject, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, id_student, id_subject, frequency, status, created_at, updated_at FROM students_subjects")
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	defer core.DbClose(rows)

	var students []domain.StudentSubject
	for rows.Next() {
		student, err := scanDefaultStudentSubject(rows)

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

func (r *Repository) FetchByStudentSubjectId(ctx context.Context, idStudent int) ([]domain.StudentSubjectWithSubject, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT ss.id,
		       ss.id_student,
		       ss.id_subject,
		       ss.frequency,
			   ss.status,
		       ss.created_at,
		       ss.updated_at,
			   su.name
		FROM students_subjects ss
				 JOIN subjects su on su.id = ss.id_subject
		WHERE ss.id_student=?`, idStudent)
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	defer core.DbClose(rows)

	var students []domain.StudentSubjectWithSubject
	for rows.Next() {
		var ss domain.StudentSubjectWithSubject
		err := rows.Scan(
			 &ss.Id,
			 &ss.IdSubject,
			 &ss.IdSubject,
			 &ss.Frequency,
			 &ss.Status,
			 &ss.CreatedAt,
			 &ss.UpdatedAt,
			 &ss.Subject.Name,
		)
		if err != nil {
			return nil, core.NewError(nil, err, 0)
		}
		ss.Subject.Id = ss.IdSubject

		students = append(students, ss)
	}

	if len(students) == 0 {
		return nil, core.NotFoundError(nil, "Student not found")
	}

	return students, err
}

func (r *Repository) Insert(ctx context.Context, std *domain.StudentSubject) (int, error) {
	res, err := r.db.ExecContext(ctx, "INSERT INTO students_subjects (id_student, id_subject, frequency, status) VALUES (?, ?, ?, ?)",
		std.IdStudent, std.IdSubject, std.Frequency, std.Status)
	if err != nil {
		return 0, core.NewError(nil, err, 0)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, core.WrapError(err)
	}

	return int(id), err
}
