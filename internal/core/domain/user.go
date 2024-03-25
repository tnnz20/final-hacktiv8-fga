package domain

import "time"

type User struct {
	ID              int
	Username        string
	Email           string
	Password        string
	Age             int
	ProfileImageURL string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CreateUserRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	Age             int    `json:"age" validate:"required,number,gt=8"`
	ProfileImageURL string `json:"profile_image_url" validate:"required,url"`
}

type CreateUserResponse struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Age             int    `json:"age "`
	ProfileImageURL string `json:"profile_image_url"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	ID              int    `json:"id"`
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Age             int    `json:"age" validate:"required,gt=8"`
	ProfileImageURL string `json:"profile_image_url" validate:"required,url"`
}

type UpdateUserResponse struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Age             int    `json:"age"`
	ProfileImageURL string `json:"profile_image_url"`
}
