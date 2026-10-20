package service

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/fisdemire/pickup_point_service/internal/domain"
)

type UserStore interface {
	CreateUser(ctx context.Context, email, passwordHash string, role domain.Role) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type TokenIssuer interface {
	Issue(userID uuid.UUID, role domain.Role) (string, error)
}

var (
	dummyEmployeeID  = uuid.MustParse("11111111-1111-1111-1111-111111111111")
	dummyModeratorID = uuid.MustParse("22222222-2222-2222-2222-222222222222")
)

type Auth struct {
	users  UserStore
	tokens TokenIssuer
}

func NewAuth(users UserStore, tokens TokenIssuer) *Auth { return &Auth{users: users, tokens: tokens} }

func (s *Auth) DummyLogin(role string) (string, error) {
	r := domain.Role(role)
	if !r.Valid() {
		return "", domain.NewValidation("role must be employee or moderator")
	}
	id := dummyEmployeeID
	if r == domain.RoleModerator {
		id = dummyModeratorID
	}
	return s.tokens.Issue(id, r)
}

func (s *Auth) Register(ctx context.Context, email, password, role string) (domain.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if addr, err := mail.ParseAddress(email); err != nil || addr.Address != email {
		return domain.User{}, domain.NewValidation("invalid email")
	}
	if l := len(password); l == 0 || l > 72 {
		return domain.User{}, domain.NewValidation("password must be 1..72 bytes")
	}
	r := domain.Role(role)
	if !r.Valid() {
		return domain.User{}, domain.NewValidation("role must be employee or moderator")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	return s.users.CreateUser(ctx, email, string(hash), r)
}

func (s *Auth) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.users.GetUserByEmail(ctx, strings.ToLower(strings.TrimSpace(email)))
	if errors.Is(err, domain.ErrUserNotFound) {
		return "", domain.ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return "", domain.ErrInvalidCredentials
	}
	return s.tokens.Issue(u.ID, u.Role)
}
