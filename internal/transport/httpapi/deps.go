package httpapi

import (
	"context"

	"github.com/fisdemire/pickup_point_service/internal/domain"
)

type AuthService interface {
	DummyLogin(role string) (string, error)
	Register(ctx context.Context, email, password, role string) (domain.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
