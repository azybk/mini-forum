package memberships

import (
	"context"
	"errors"
	"log"

	"github.com/azybk/mini-forum/internal/model/memberships"
	"github.com/azybk/mini-forum/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Fatal("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email is not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("email or password is invalid")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}