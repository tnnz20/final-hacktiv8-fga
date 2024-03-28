package port

import (
	"context"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type SocialMediaRepository interface {
	Create(ctx context.Context, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
	Update(ctx context.Context, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
	Delete(ctx context.Context, socialMediaId int) error
	FindSocialMediaByID(ctx context.Context, socialMediaId int) (*domain.SocialMedia, error)
	FindSocialMedias(ctx context.Context) (*[]domain.SocialMedia, error)
}

type SocialMediaService interface {
	Create(ctx context.Context, req *domain.CreateSocialMediaRequest, userId int) (*domain.CreateSocialMediaResponse, error)
	Update(ctx context.Context, req *domain.UpdateSocialMediaRequest, socialMediaId, userId int) (*domain.UpdateSocialMediaResponse, error)
	Delete(ctx context.Context, socialMediaId, userId int) error
	GetSocialMediaByID(ctx context.Context, socialMediaId int) (*domain.GetSocialMediaResponse, error)
	GetSocialMedias(ctx context.Context) (*[]domain.GetSocialMediaResponse, error)
}
