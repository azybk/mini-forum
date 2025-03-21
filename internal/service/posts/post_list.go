package posts

import (
	"context"
	"log"

	"github.com/azybk/mini-forum/internal/model/posts"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := limit * (pageIndex - 1)

	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Println("error get all post from database")
		return response, err
	}

	return response, nil
}