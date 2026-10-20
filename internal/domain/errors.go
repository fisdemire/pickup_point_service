package domain

import "errors"

var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrUserNotFound        = errors.New("user not found")
	ErrEmailTaken          = errors.New("email already registered")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrPVZNotFound         = errors.New("pvz not found")
	ErrOpenReceptionExists = errors.New("previous reception is not closed")
	ErrNoOpenReception     = errors.New("no open reception")
	ErrNoProducts          = errors.New("no products to delete")
)

type ValidationError struct{ Msg string }

func (e *ValidationError) Error() string { return e.Msg }

func NewValidation(msg string) error { return &ValidationError{Msg: msg} }
