package domain

type Comment struct {
	ID      int
	Message string
	PhotoID int
	UserID  int
}

type CreateCommentRequest struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}

type CreateCommentResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
}

type GetCommentResponse struct {
	ID      int         `json:"id"`
	Message string      `json:"message"`
	PhotoID int         `json:"photo_id"`
	UserID  int         `json:"user_id"`
	User    UserDetail  `json:"user"`
	Photo   PhotoDetail `json:"photo"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required"`
}

type UpdateCommentResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
}
