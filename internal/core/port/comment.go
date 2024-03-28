package port

import (
	"context"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error)
	Update(ctx context.Context, comment *domain.Comment) (*domain.Comment, error)
	Delete(ctx context.Context, commentId int) error
	FindCommentByID(ctx context.Context, commentId int) (*domain.Comment, error)
	FindComments(ctx context.Context) (*[]domain.Comment, error)
}

type CommentService interface {
	Create(ctx context.Context, req *domain.CreateCommentRequest, userId int) (*domain.CreateCommentResponse, error)
	Update(ctx context.Context, req *domain.UpdateCommentRequest, commentId, userId int) (*domain.UpdateCommentResponse, error)
	Delete(ctx context.Context, commentId, userId int) error
	GetCommentByID(ctx context.Context, commentId int) (*domain.GetCommentResponse, error)
	GetComments(ctx context.Context) (*[]domain.GetCommentResponse, error)
}
