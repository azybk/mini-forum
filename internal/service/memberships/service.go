package memberships

import (
	"context"

	"github.com/azybk/mini-forum/internal/configs"
	"github.com/azybk/mini-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg: cfg,
		membershipRepo: membershipRepo,
	}
}
