package posts

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/azybk/mini-forum/internal/model/posts"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Println("failed to create comment to reposiroty")
		return err
	}
	return nil
}
