package domain

type Photo struct {
	ID       int
	Title    string
	Caption  string
	PhotoURL string
	UserID   int
}

type CreatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

type CreatePhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type userPhoto struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type GetPhoto struct {
	ID       int       `json:"id"`
	Caption  string    `json:"caption"`
	Title    string    `json:"title"`
	PhotoURL string    `json:"photo_url"`
	UserID   int       `json:"user_id"`
	User     userPhoto `json:"user"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

type UpdatePhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}
