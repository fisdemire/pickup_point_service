package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/fisdemire/pickup_point_service/internal/domain"
)

type Identity struct {
	UserID uuid.UUID
	Role   domain.Role
}

type claims struct {
	Role domain.Role `json:"role"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{secret: []byte(secret), ttl: ttl}
}

func (m *TokenManager) Issue(userID uuid.UUID, role domain.Role) (string, error) {
	now := time.Now()
	c := claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(m.secret)
}

func (m *TokenManager) Parse(raw string) (Identity, error) {
	var c claims
	_, err := jwt.ParseWithClaims(raw, &c,
		func(*jwt.Token) (any, error) { return m.secret, nil },
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return Identity{}, domain.ErrUnauthorized
	}
	id, err := uuid.Parse(c.Subject)
	if err != nil || !c.Role.Valid() {
		return Identity{}, domain.ErrUnauthorized
	}
	return Identity{UserID: id, Role: c.Role}, nil
}
