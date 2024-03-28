package service

import (
	"context"
	"time"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/port"
)

type PhotoService struct {
	PhotoRepository port.PhotoRepository
	UserRepository  port.UserRepository
	timeout         time.Duration
}

func NewPhotoService(photoRepo port.PhotoRepository, userRepo port.UserRepository) *PhotoService {
	return &PhotoService{
		PhotoRepository: photoRepo,
		UserRepository:  userRepo,
		timeout:         time.Duration(3) * time.Second,
	}
}

func (s *PhotoService) Create(ctx context.Context, req *domain.CreatePhotoRequest, userId int) (*domain.CreatePhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photo := &domain.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   userId,
	}

	createdPhoto, err := s.PhotoRepository.Create(ctx, photo)
	if err != nil {
		return nil, err
	}

	res := &domain.CreatePhotoResponse{
		ID:       createdPhoto.ID,
		Title:    createdPhoto.Title,
		Caption:  createdPhoto.Caption,
		PhotoURL: createdPhoto.PhotoURL,
		UserID:   createdPhoto.UserID,
	}

	return res, nil
}

func (s *PhotoService) Update(ctx context.Context, req *domain.UpdatePhotoRequest, photoId, userId int) (*domain.UpdatePhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photo, err := s.PhotoRepository.FindPhotoByID(ctx, photoId)
	if err != nil {
		return nil, err
	}

	if photo.UserID != userId {
		return nil, domain.ErrUnauthorized
	}

	photo.Title = req.Title
	photo.Caption = req.Caption
	photo.PhotoURL = req.PhotoURL

	updatedPhoto, err := s.PhotoRepository.Update(ctx, photo)
	if err != nil {
		return nil, err
	}

	res := &domain.UpdatePhotoResponse{
		ID:       updatedPhoto.ID,
		Title:    updatedPhoto.Title,
		Caption:  updatedPhoto.Caption,
		PhotoURL: updatedPhoto.PhotoURL,
		UserID:   updatedPhoto.UserID,
	}

	return res, nil
}

func (s *PhotoService) Delete(ctx context.Context, photoId, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photo, err := s.PhotoRepository.FindPhotoByID(ctx, photoId)
	if err != nil {
		return err
	}

	if photo.UserID != userId {
		return domain.ErrUnauthorized
	}

	if err := s.PhotoRepository.Delete(ctx, photo.ID); err != nil {
		return err
	}

	return nil
}

func (s *PhotoService) GetPhotos(ctx context.Context) (*[]domain.GetPhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photos, err := s.PhotoRepository.FindPhotos(ctx)
	if err != nil {
		return nil, err
	}

	var responses []domain.GetPhotoResponse
	for _, photo := range *photos {
		user, err := s.UserRepository.FindUserById(ctx, photo.UserID)
		if err != nil {
			return nil, err
		}

		userDetail := domain.UserDetail{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		res := domain.GetPhotoResponse{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserID:   photo.UserID,
			User:     userDetail,
		}

		responses = append(responses, res)

	}

	return &responses, nil
}

func (s *PhotoService) GetPhotoByID(ctx context.Context, photoId int) (*domain.GetPhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photo, err := s.PhotoRepository.FindPhotoByID(ctx, photoId)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindUserById(ctx, photo.UserID)
	if err != nil {
		return nil, err
	}

	userDetail := domain.UserDetail{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	res := &domain.GetPhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
		User:     userDetail,
	}

	return res, nil
}
