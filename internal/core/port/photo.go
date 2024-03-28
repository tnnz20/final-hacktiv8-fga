package port

import (
	"context"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type PhotoRepository interface {
	Create(ctx context.Context, photo *domain.Photo) (*domain.Photo, error)
	Update(ctx context.Context, photo *domain.Photo) (*domain.Photo, error)
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*domain.GetPhoto, error)
	FindAll(ctx context.Context) (*[]domain.GetPhoto, error)
	FindByUserID(ctx context.Context, userId int) (*domain.Photo, error)
}

type PhotoService interface {
	Create(ctx context.Context, req *domain.CreatePhotoRequest, userId int) (*domain.CreatePhotoResponse, error)
	Update(ctx context.Context, req *domain.UpdatePhotoRequest, photoId, userId int) (*domain.UpdatePhotoResponse, error)
	Delete(ctx context.Context, photoId, userId int) error
	GetByID(ctx context.Context, id int) (*domain.GetPhoto, error)
	GetAll(ctx context.Context) (*[]domain.GetPhoto, error)
}
