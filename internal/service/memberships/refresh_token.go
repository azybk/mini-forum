package memberships

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/azybk/mini-forum/internal/model/memberships"
	"github.com/azybk/mini-forum/pkg/jwt"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken , err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Println("error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	if existingRefreshToken.RefreshToken != req.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Fatal("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user is not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}