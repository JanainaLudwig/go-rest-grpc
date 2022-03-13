package subject

import (
	"context"
	"database/sql"
	"grpc-rest/core"
	"grpc-rest/domain"
)

type Repository struct {
	db *sql.DB
}

func scanDefaultSubject(rows *sql.Rows) (*domain.Subject, error) {
	var subject domain.Subject

	var created, updated sql.NullTime
	err := rows.Scan(
		&subject.Id,
		&subject.Name,
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

func (r *Repository) FetchAll(ctx context.Context) ([]domain.Subject, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, created_at, updated_at FROM subjects")
	if err != nil {
		return nil, core.NewError(nil, err, 0)
	}

	defer core.DbClose(rows)

	var subjects []domain.Subject
	for rows.Next() {
		subject, err := scanDefaultSubject(rows)

		if err != nil {
			return nil, core.NewError(nil, err, 0)
		}

		subjects = append(subjects, *subject)
	}

	if len(subjects) == 0 {
		return nil, core.NotFoundError(nil, "Subject not found")
	}

	return subjects, err
}

func (r *Repository) Insert(ctx context.Context, std *domain.Subject) (int, error) {
	res, err := r.db.ExecContext(ctx, "INSERT INTO subjects (name) VALUES (?)",
		std.Name)
	if err != nil {
		return 0, core.NewError(nil, err, 0)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, core.WrapError(err)
	}

	return int(id), err
}
