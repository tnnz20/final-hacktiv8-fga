package service

import (
	"context"
	"time"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/port"
)

type SocialMediaService struct {
	SocialMediaRepository port.SocialMediaRepository
	UserRepository        port.UserRepository
	timeout               time.Duration
}

func NewSocialMediaService(socialMediaRepo port.SocialMediaRepository, userRepo port.UserRepository) *SocialMediaService {
	return &SocialMediaService{
		SocialMediaRepository: socialMediaRepo,
		UserRepository:        userRepo,
		timeout:               time.Duration(3) * time.Second,
	}
}

func (s *SocialMediaService) Create(ctx context.Context, req *domain.CreateSocialMediaRequest, userId int) (*domain.CreateSocialMediaResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	socialMedia := &domain.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         userId,
	}

	createdSocialMedia, err := s.SocialMediaRepository.Create(ctx, socialMedia)
	if err != nil {
		return nil, err
	}

	res := &domain.CreateSocialMediaResponse{
		ID:             createdSocialMedia.ID,
		Name:           createdSocialMedia.Name,
		SocialMediaURL: createdSocialMedia.SocialMediaURL,
		UserID:         createdSocialMedia.UserID,
	}

	return res, nil
}

func (s *SocialMediaService) Update(ctx context.Context, req *domain.UpdateSocialMediaRequest, socialMediaId, userId int) (*domain.UpdateSocialMediaResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	socialMedia, err := s.SocialMediaRepository.FindSocialMediaByID(ctx, socialMediaId)
	if err != nil {
		return nil, err
	}

	if socialMedia.UserID != userId {
		return nil, domain.ErrUnauthorized
	}

	socialMedia.Name = req.Name
	socialMedia.SocialMediaURL = req.SocialMediaURL

	updatedSocialMedia, err := s.SocialMediaRepository.Update(ctx, socialMedia)
	if err != nil {
		return nil, err
	}

	res := &domain.UpdateSocialMediaResponse{
		ID:             updatedSocialMedia.ID,
		Name:           updatedSocialMedia.Name,
		SocialMediaURL: updatedSocialMedia.SocialMediaURL,
		UserID:         updatedSocialMedia.UserID,
	}

	return res, nil
}

func (s *SocialMediaService) Delete(ctx context.Context, socialMediaId, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	socialMedia, err := s.SocialMediaRepository.FindSocialMediaByID(ctx, socialMediaId)
	if err != nil {
		return err
	}

	if socialMedia.UserID != userId {
		return domain.ErrUnauthorized
	}

	err = s.SocialMediaRepository.Delete(ctx, socialMedia.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SocialMediaService) GetSocialMedias(ctx context.Context) (*[]domain.GetSocialMediaResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	socialMedias, err := s.SocialMediaRepository.FindSocialMedias(ctx)
	if err != nil {
		return nil, err
	}

	var responses []domain.GetSocialMediaResponse
	for _, socialMedia := range *socialMedias {
		user, err := s.UserRepository.FindUserById(ctx, socialMedia.UserID)
		if err != nil {
			return nil, err
		}

		userDetail := domain.UserDetail{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		res := domain.GetSocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         socialMedia.UserID,
			User:           userDetail,
		}

		responses = append(responses, res)

	}

	return &responses, nil
}

func (s *SocialMediaService) GetSocialMediaByID(ctx context.Context, socialMediaId int) (*domain.GetSocialMediaResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	socialMedia, err := s.SocialMediaRepository.FindSocialMediaByID(ctx, socialMediaId)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindUserById(ctx, socialMedia.UserID)
	if err != nil {
		return nil, err
	}

	userDetail := domain.UserDetail{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	res := &domain.GetSocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		User:           userDetail,
	}

	return res, nil
}
