package posts

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/azybk/mini-forum/internal/model/posts"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID: postID,
		UserID: userID,
		IsLiked: request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Println("error get user activity to database")
		return err
	}

	if userActivity == nil {
		// create user activity
		if !request.IsLiked {
			return errors.New("you haven't like before")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)

	} else {
		// update user activity
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Println("error create or update user activity to database")
		return err
	}
	
	return nil
}