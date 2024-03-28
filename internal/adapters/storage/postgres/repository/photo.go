package repository

import (
	"context"
	"database/sql"

	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/storage/postgres"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type PhotoRepository struct {
	Db postgres.DBTX
}

func NewPhotoRepository(db postgres.DBTX) *PhotoRepository {
	return &PhotoRepository{
		Db: db,
	}
}

func (r PhotoRepository) Create(ctx context.Context, photo *domain.Photo) (*domain.Photo, error) {
	var id int

	query := `
		INSERT INTO photos (title, caption, photo_url, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.Db.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoURL, photo.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}

	photo.ID = id

	return photo, nil
}

func (r PhotoRepository) Update(ctx context.Context, photo *domain.Photo) (*domain.Photo, error) {
	query := `
		UPDATE photos
		SET title = $1, caption = $2, photo_url = $3
		WHERE id = $4
		RETURNING id, title, caption, photo_url, user_id
	`

	err := r.Db.QueryRowContext(ctx, query, photo.Title, photo.Caption, photo.PhotoURL,
		photo.ID).Scan(&photo.ID, &photo.Title, &photo.Caption,
		&photo.PhotoURL, &photo.UserID)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (r PhotoRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM photos
		WHERE id = $1
	`
	_, err := r.Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r PhotoRepository) FindByID(ctx context.Context, id int) (*domain.GetPhoto, error) {
	var photo domain.GetPhoto

	query := `
		SELECT id, title, caption, photo_url, user_id
		FROM photos
		WHERE id = $1
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&photo.ID, &photo.Title, &photo.Caption,
		&photo.PhotoURL, &photo.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrPhotoNotFound
		}
		return nil, err
	}

	query = `
		SELECT id, email, username 
		FROM users 
		WHERE id = $1
	`

	err = r.Db.QueryRowContext(ctx, query, photo.UserID).Scan(&photo.User.ID,
		&photo.User.Email, &photo.User.Username)
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

func (r PhotoRepository) FindAll(ctx context.Context) (*[]domain.GetPhoto, error) {
	var photos []domain.GetPhoto

	query := `
		SELECT id, title, caption, photo_url, user_id
		FROM photos
	`

	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		if rows.Err() == sql.ErrNoRows {
			return nil, domain.ErrPhotosEmpty
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo domain.GetPhoto
		err := rows.Scan(&photo.ID, &photo.Title, &photo.Caption,
			&photo.PhotoURL, &photo.UserID)
		if err != nil {
			return nil, err
		}

		query := `SELECT id, email, username FROM users WHERE id = $1`
		err = r.Db.QueryRowContext(ctx, query, photo.UserID).Scan(&photo.User.ID, &photo.User.Email, &photo.User.Username)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return &photos, nil
}

func (r PhotoRepository) FindByUserID(ctx context.Context, userId int) (*domain.Photo, error) {
	var photo domain.Photo

	query := `
		SELECT id, title, caption, photo_url, user_id
		FROM photos
		WHERE user_id = $1
	`

	err := r.Db.QueryRowContext(ctx, query, userId).Scan(&photo.ID, &photo.Title, &photo.Caption,
		&photo.PhotoURL, &photo.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrPhotoNotFound
		}
		return nil, err
	}

	return &photo, nil
}
