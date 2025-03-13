package posts

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/azybk/mini-forum/internal/model/posts"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHastaghs := strings.Join(req.PostHashtags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID: userID,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHashtags: postHastaghs,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Println("error Create Post to repository")
		return err
	}

	return nil
}