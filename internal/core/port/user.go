package port

import (
	"context"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id int) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserById(ctx context.Context, id int) (*domain.User, error)
}

type UserService interface {
	Register(ctx context.Context, req *domain.CreateUserRequest) (*domain.CreateUserResponse, error)
	Login(ctx context.Context, req *domain.LoginUserRequest) (*domain.LoginUserResponse, error)
	Update(ctx context.Context, req *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error)
	Delete(ctx context.Context, id int) error
}
