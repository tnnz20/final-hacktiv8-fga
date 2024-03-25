package repository

import (
	"context"

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
	err := r.Db.QueryRowContext(ctx, query, photo.Title,
		photo.Caption, photo.PhotoURL, photo.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}

	p := &domain.Photo{
		ID:       id,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}

	return p, nil
}

func (r PhotoRepository) Update(ctx context.Context, photo *domain.Photo) (*domain.Photo, error) {
	query := `
		UPDATE photos
		SET title = $1, caption = $2, photo_url = $3
		WHERE id = $4
	`
	_, err := r.Db.ExecContext(ctx, query, photo.Title, photo.Caption, photo.PhotoURL, photo.ID)
	if err != nil {
		return nil, err
	}

	return photo, nil
}
