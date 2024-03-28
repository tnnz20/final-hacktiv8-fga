package domain

type SocialMedia struct {
	ID             int
	Name           string
	SocialMediaURL string
	UserID         int
}

type CreateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

type CreateSocialMediaResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
}

type GetSocialMediaResponse struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         int        `json:"user_id"`
	User           UserDetail `json:"user"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

type UpdateSocialMediaResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
}
