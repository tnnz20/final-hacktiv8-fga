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

func (s *PhotoService) GetAll(ctx context.Context) (*[]domain.GetPhoto, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photos, err := s.PhotoRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, photo := range *photos {
		user, err := s.UserRepository.GetUserById(ctx, photo.UserID)
		if err != nil {
			return nil, err
		}

		(*photos)[i].User.ID = user.ID
		(*photos)[i].User.Username = user.Username
		(*photos)[i].User.Email = user.Email
	}

	return photos, nil
}

func (s *PhotoService) GetByID(ctx context.Context, id int) (*domain.GetPhoto, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	photo, err := s.PhotoRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.GetUserById(ctx, photo.UserID)
	if err != nil {
		return nil, err
	}

	photo.User.ID = user.ID
	photo.User.Username = user.Username
	photo.User.Email = user.Email

	return photo, nil
}

func (s *PhotoService) Update(ctx context.Context, req *domain.UpdatePhotoRequest, photoId, userId int) (*domain.UpdatePhotoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	checkPhoto, err := s.PhotoRepository.FindByID(ctx, photoId)
	if err != nil {
		return nil, err
	} else if checkPhoto.UserID != userId {
		return nil, domain.ErrUnauthorized
	}

	photo := &domain.Photo{
		ID:       photoId,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   userId,
	}

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

	checkPhoto, err := s.PhotoRepository.FindByID(ctx, photoId)
	if err != nil {
		return err
	} else if checkPhoto.UserID != userId {
		return domain.ErrUnauthorized
	}

	if err := s.PhotoRepository.Delete(ctx, photoId); err != nil {
		return err
	}

	return nil
}
