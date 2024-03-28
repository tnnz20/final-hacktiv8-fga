package domain

type Photo struct {
	ID       int
	Caption  string
	Title    string
	PhotoURL string
	UserID   int
}

type PhotoDetail struct {
	ID       int    `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}
type CreatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

type CreatePhotoResponse struct {
	ID       int    `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type GetPhotoResponse struct {
	ID       int        `json:"id"`
	Caption  string     `json:"caption"`
	Title    string     `json:"title"`
	PhotoURL string     `json:"photo_url"`
	UserID   int        `json:"user_id"`
	User     UserDetail `json:"user"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

type UpdatePhotoResponse struct {
	ID       int    `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}
