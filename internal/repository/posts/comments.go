package posts

import (
	"context"

	"github.com/azybk/mini-forum/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := "INSERT INTO comments(post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES(?, ?, ?, ?, ?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCommentsByPostID(ctx context.Context, postID int64) ([]posts.Comment, error) {
	query := "SELECT c.id, c.user_id, u.username, c.comment_content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id = ?"

	row, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	response := make([]posts.Comment, 0)
	for row.Next() {
		var(
			comment posts.Comment
		)
		err := row.Scan(&comment.ID, &comment.UserID, &comment.Username, &comment.CommentContent)
		if err != nil {
			return nil, err
		}

		response = append(response, posts.Comment{
			ID: comment.ID,
			UserID: comment.UserID,
			Username: comment.Username,
			CommentContent: comment.CommentContent,
		})
	}

	return response, nil
}