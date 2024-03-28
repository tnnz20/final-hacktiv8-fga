package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/port"
	"github.com/tnnz20/final-hacktiv8-fga/pkg/utils"
)

type UserService struct {
	UserRepository port.UserRepository
	timeout        time.Duration
	jwtKey         *string
}

func NewUserService(repository port.UserRepository, key *string) *UserService {
	return &UserService{
		UserRepository: repository,
		timeout:        time.Duration(3) * time.Second,
		jwtKey:         key,
	}
}

func (s *UserService) Register(ctx context.Context, req *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.UserRepository.GetUserByEmail(ctx, req.Email)
	if err == nil {
		if user.Email == req.Email {
			return nil, domain.ErrEmailExist
		} else if user.Username == req.Username {
			return nil, domain.ErrUsernameExist
		}
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	userReq := &domain.User{
		Username:        req.Username,
		Email:           req.Email,
		Password:        hashedPassword,
		Age:             req.Age,
		ProfileImageURL: req.ProfileImageURL,
	}

	userRes, err := s.UserRepository.Create(ctx, userReq)
	if err != nil {
		return nil, err
	}

	response := &domain.CreateUserResponse{
		ID:              userRes.ID,
		Username:        userRes.Username,
		Email:           userRes.Email,
		Password:        userRes.Password,
		Age:             userRes.Age,
		ProfileImageURL: userRes.ProfileImageURL,
	}

	return response, nil
}

func (s *UserService) Login(ctx context.Context, req *domain.LoginUserRequest) (*domain.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.UserRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err = utils.ComparePassword(req.Password, user.Password); err != nil {
		return nil, domain.ErrWrongPassword
	}

	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	token, err := t.SignedString([]byte(*s.jwtKey))
	if err != nil {
		return nil, err
	}

	response := &domain.LoginUserResponse{
		Token: token,
	}

	return response, nil
}

func (s *UserService) Update(ctx context.Context, req *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.UserRepository.GetUserById(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.ProfileImageURL == "" {
		req.ProfileImageURL = user.ProfileImageURL
	}

	userReq := &domain.User{
		ID:              req.ID,
		Username:        req.Username,
		Email:           req.Email,
		Age:             req.Age,
		ProfileImageURL: req.ProfileImageURL,
	}

	userRes, err := s.UserRepository.Update(ctx, userReq)
	if err != nil {
		return nil, err
	}

	response := &domain.UpdateUserResponse{
		ID:              userRes.ID,
		Username:        userRes.Username,
		Email:           userRes.Email,
		Age:             userRes.Age,
		ProfileImageURL: userRes.ProfileImageURL,
	}

	return response, nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.UserRepository.GetUserById(ctx, id); err != nil {
		return err
	}

	err := s.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
