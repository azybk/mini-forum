package memberships

import (
	"context"

	"github.com/azybk/mini-forum/internal/middleware"
	"github.com/azybk/mini-forum/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.RefreshToken)
}
