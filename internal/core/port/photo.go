package port

import (
	"context"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type PhotoRepository interface {
	Create(ctx context.Context, photo *domain.Photo) (*domain.Photo, error)
	Update(ctx context.Context, photo *domain.Photo) (*domain.Photo, error)
	Delete(ctx context.Context, photoId int) error
	FindPhotoByID(ctx context.Context, photoId int) (*domain.GetPhotoResponse, error)
	FindPhotos(ctx context.Context) (*[]domain.GetPhotoResponse, error)
}

type PhotoService interface {
	Create(ctx context.Context, req *domain.CreatePhotoRequest, userId int) (*domain.CreatePhotoResponse, error)
	Update(ctx context.Context, req *domain.UpdatePhotoRequest, photoId, userId int) (*domain.UpdatePhotoResponse, error)
	Delete(ctx context.Context, photoId, userId int) error
	GetPhotoByID(ctx context.Context, photoId int) (*domain.GetPhotoResponse, error)
	GetPhotos(ctx context.Context) (*[]domain.GetPhotoResponse, error)
}
