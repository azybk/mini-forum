package memberships

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/azybk/mini-forum/internal/model/memberships"
	"github.com/azybk/mini-forum/pkg/jwt"
	tokenUtil "github.com/azybk/mini-forum/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Fatal("failed to get user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email is not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", errors.New("email or password is invalid")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	existRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Println("error get latest refresh token from database")
		return "", "", err
	}

	if existRefreshToken != nil {
		return token, existRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID: user.ID,
		RefreshToken: refreshToken,
		ExpiredAt: time.Now().Add(10 * 24 * time.Hour),
	})

	if err != nil {
		log.Println("error insert refresh token to database")
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}