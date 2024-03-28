package repository

import (
	"context"
	"database/sql"

	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/storage/postgres"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
)

type CommentRepository struct {
	Db postgres.DBTX
}

func NewCommentRepository(db postgres.DBTX) *CommentRepository {
	return &CommentRepository{
		Db: db,
	}
}

func (r *CommentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	query := `
		INSERT INTO comments (message, photo_id, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.Db.QueryRowContext(ctx, query, comment.Message, comment.PhotoID, comment.UserID).Scan(&comment.ID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) Update(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	query := `
		UPDATE comments 
		SET message = $1 
		WHERE id = $2 
		RETURNING photo_id, user_id
	`
	err := r.Db.QueryRowContext(ctx, query, comment.Message, comment.ID).Scan(&comment.PhotoID, &comment.UserID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) Delete(ctx context.Context, commentId int) error {
	query := `
		DELETE FROM comments
		WHERE id = $1
	`

	res, err := r.Db.ExecContext(ctx, query, commentId)
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		if row == 0 {
			return err
		}
		return err
	}

	return nil
}

func (r *CommentRepository) FindComments(ctx context.Context) (*[]domain.Comment, error) {
	query := `SELECT id, message, photo_id, user_id FROM comments`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrCommentsEmpty
		}
		return nil, err
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		var comment domain.Comment
		err := rows.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return &comments, nil
}

func (r *CommentRepository) FindCommentByID(ctx context.Context, commentId int) (*domain.Comment, error) {
	query := `SELECT id, message, photo_id, user_id FROM comments WHERE id = $1`
	row := r.Db.QueryRowContext(ctx, query, commentId)

	var comment domain.Comment
	err := row.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrCommentNotFound
		}
		return nil, err
	}

	return &comment, nil
}
