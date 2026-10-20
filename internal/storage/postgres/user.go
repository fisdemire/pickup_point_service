package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/fisdemire/pickup_point_service/internal/domain"
)

func (s *Store) CreateUser(ctx context.Context, email, hash string, role domain.Role) (domain.User, error) {
	u := domain.User{Email: email, Role: role}
	err := s.pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash, role) VALUES ($1, $2, $3) RETURNING id`,
		email, hash, role).Scan(&u.ID)
	if violates(err, pgUniqueViolation, "") {
		return domain.User{}, domain.ErrEmailTaken
	}
	return u, err
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var u domain.User
	err := s.pool.QueryRow(ctx,
		`SELECT id, email, password_hash, role FROM users WHERE email = $1`, email).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, domain.ErrUserNotFound
	}
	return u, err
}
