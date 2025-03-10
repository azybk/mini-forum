package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/azybk/mini-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("email or username already exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Username,
		UpdatedBy: req.Username,
		Username:  req.Username,
	}

	return s.membershipRepo.CreateUser(ctx, model)

}
