package repository

import (
	"context"
	"database/sql"

	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/storage/postgres"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type SocialMediaRepository struct {
	Db postgres.DBTX
}

func NewSocialMediaRepository(db postgres.DBTX) *SocialMediaRepository {
	return &SocialMediaRepository{
		Db: db,
	}
}

func (r SocialMediaRepository) Create(ctx context.Context, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	var id int

	query := `
		INSERT INTO social_medias (name, social_media_url, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.Db.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaURL, socialMedia.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}

	socialMedia.ID = id

	return socialMedia, nil
}

func (r SocialMediaRepository) Update(ctx context.Context, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	query := `
		UPDATE social_medias
		SET name = $1, social_media_url = $2
		WHERE id = $3
		RETURNING id, name, social_media_url, user_id
	`

	err := r.Db.QueryRowContext(ctx, query, socialMedia.Name, socialMedia.SocialMediaURL,
		socialMedia.ID).Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserID)
	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (r SocialMediaRepository) Delete(ctx context.Context, socialMediaId int) error {
	query := `
		DELETE FROM social_medias
		WHERE id = $1
	`

	res, err := r.Db.ExecContext(ctx, query, socialMediaId)
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

func (r SocialMediaRepository) FindSocialMediaByID(ctx context.Context, socialMediaId int) (*domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia

	query := `
		SELECT id, name, social_media_url, user_id
		FROM social_medias
		WHERE id = $1
	`
	err := r.Db.QueryRowContext(ctx, query, socialMediaId).Scan(&socialMedia.ID, &socialMedia.Name,
		&socialMedia.SocialMediaURL, &socialMedia.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrSocialMediaNotFound
		}
		return nil, err
	}

	return &socialMedia, nil
}

func (r SocialMediaRepository) FindSocialMedias(ctx context.Context) (*[]domain.SocialMedia, error) {
	var socialMedias []domain.SocialMedia

	query := `
		SELECT id, name, social_media_url, user_id
		FROM social_medias
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
		var socialMedia domain.SocialMedia
		err := rows.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserID)
		if err != nil {
			return nil, err
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	return &socialMedias, nil
}
