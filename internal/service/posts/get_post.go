package posts

import (
	"context"
	"log"

	"github.com/azybk/mini-forum/internal/model/posts"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Println("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Println("error get like count to database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		log.Println("error get comments to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID: postID,
			UserID: postDetail.UserID,
			Username: postDetail.Username,
			PostTitle: postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked: postDetail.IsLiked,
		}, 
		LikeCount: likeCount,  
		Comments: comments,   
	}, nil
}