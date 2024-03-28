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

func (r PhotoRepository) Delete(ctx context.Context, photoId int) error {
	query := `
		DELETE FROM photos
		WHERE id = $1
	`
	res, err := r.Db.ExecContext(ctx, query, photoId)
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		if row == 0 {
			return err
		}
		return err
	}

	return nil
}

func (r PhotoRepository) FindPhotoByID(ctx context.Context, photoId int) (*domain.GetPhotoResponse, error) {
	var photo domain.GetPhotoResponse

	query := `
		SELECT id, title, caption, photo_url, user_id
		FROM photos
		WHERE id = $1
	`
	err := r.Db.QueryRowContext(ctx, query, photoId).Scan(&photo.ID, &photo.Title, &photo.Caption,
		&photo.PhotoURL, &photo.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrPhotoNotFound
		}
		return nil, err
	}

	return &photo, nil
}

func (r PhotoRepository) FindPhotos(ctx context.Context) (*[]domain.GetPhotoResponse, error) {
	var photos []domain.GetPhotoResponse

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
		var photo domain.GetPhotoResponse
		err := rows.Scan(&photo.ID, &photo.Title, &photo.Caption,
			&photo.PhotoURL, &photo.UserID)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return &photos, nil
}
