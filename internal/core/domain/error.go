package domain

import "errors"

var (
	ErrUnauthorized = errors.New("you are not authorized to change this data")
	ErrInvalidId    = errors.New("invalid id must be a number")

	ErrUserNotFound  = errors.New("user not found")
	ErrEmailExist    = errors.New("email already exists")
	ErrUsernameExist = errors.New("username already exists")
	ErrWrongPassword = errors.New("wrong password")

	ErrPhotoNotFound = errors.New("photo not found")
	ErrPhotosEmpty   = errors.New("photos still empty")

	ErrCommentNotFound = errors.New("comment not found")
	ErrCommentsEmpty   = errors.New("comments still empty")

	ErrSocialMediaNotFound = errors.New("social media not found")
	ErrSocialMediasEmpty   = errors.New("social medias still empty")
)
