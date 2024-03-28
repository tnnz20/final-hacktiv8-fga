package domain

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrPhotoNotFound = errors.New("photo not found")
	ErrPhotosEmpty   = errors.New("photos still empty")
	ErrUnauthorized  = errors.New("you are not authorized to update")
)
