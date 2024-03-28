package service

import (
	"context"
	"time"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/port"
)

type CommentService struct {
	CommentRepository port.CommentRepository
	UserRepository    port.UserRepository
	PhotoRepository   port.PhotoRepository
	timeout           time.Duration
}

func NewCommentService(commentRepo port.CommentRepository, userRepo port.UserRepository,
	photoRepo port.PhotoRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepo,
		UserRepository:    userRepo,
		PhotoRepository:   photoRepo,
		timeout:           time.Duration(3) * time.Second,
	}
}

func (s *CommentService) Create(ctx context.Context, req *domain.CreateCommentRequest, userId int) (*domain.CreateCommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// Check if photo exists
	_, err := s.PhotoRepository.FindPhotoByID(ctx, req.PhotoID)
	if err != nil {
		return nil, err
	}

	comment := &domain.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  userId,
	}

	createdComment, err := s.CommentRepository.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	res := &domain.CreateCommentResponse{
		ID:      createdComment.ID,
		Message: createdComment.Message,
		PhotoID: createdComment.PhotoID,
		UserID:  createdComment.UserID,
	}

	return res, nil
}

func (s *CommentService) Update(ctx context.Context, req *domain.UpdateCommentRequest, commentId, userId int) (*domain.UpdateCommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	comment, err := s.CommentRepository.FindCommentByID(ctx, commentId)
	if err != nil {
		return nil, err
	}

	if comment.UserID != userId {
		return nil, domain.ErrUnauthorized
	}

	comment.Message = req.Message

	updatedComment, err := s.CommentRepository.Update(ctx, comment)
	if err != nil {
		return nil, err
	}

	res := &domain.UpdateCommentResponse{
		ID:      updatedComment.ID,
		Message: updatedComment.Message,
		PhotoID: updatedComment.PhotoID,
		UserID:  updatedComment.UserID,
	}

	return res, nil
}

func (s *CommentService) Delete(ctx context.Context, commentId, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	comment, err := s.CommentRepository.FindCommentByID(ctx, commentId)
	if err != nil {
		return err
	}

	if comment.UserID != userId {
		return domain.ErrUnauthorized
	}

	err = s.CommentRepository.Delete(ctx, comment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetComments(ctx context.Context) (*[]domain.GetCommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	comments, err := s.CommentRepository.FindComments(ctx)
	if err != nil {
		return nil, err
	}

	var responses []domain.GetCommentResponse
	for i, comment := range *comments {

		user, err := s.UserRepository.FindUserById(ctx, comment.UserID)
		if err != nil {
			return nil, err
		}

		userDetail := domain.UserDetail{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		photo, err := s.PhotoRepository.FindPhotoByID(ctx, comment.PhotoID)
		if err != nil {
			return nil, err
		}

		photoDetail := domain.PhotoDetail{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserID:   photo.UserID,
		}

		res := domain.GetCommentResponse{
			ID:      (*comments)[i].ID,
			Message: (*comments)[i].Message,
			PhotoID: (*comments)[i].PhotoID,
			UserID:  (*comments)[i].UserID,
			User:    userDetail,
			Photo:   photoDetail,
		}

		responses = append(responses, res)
	}

	return &responses, nil
}

func (s *CommentService) GetCommentByID(ctx context.Context, commentId int) (*domain.GetCommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	comment, err := s.CommentRepository.FindCommentByID(ctx, commentId)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindUserById(ctx, comment.UserID)
	if err != nil {
		return nil, err
	}

	userDetail := domain.UserDetail{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	photo, err := s.PhotoRepository.FindPhotoByID(ctx, comment.PhotoID)
	if err != nil {
		return nil, err
	}

	photoDetail := domain.PhotoDetail{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}

	res := domain.GetCommentResponse{
		ID:      comment.ID,
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID:  comment.UserID,
		User:    userDetail,
		Photo:   photoDetail,
	}

	return &res, nil
}
