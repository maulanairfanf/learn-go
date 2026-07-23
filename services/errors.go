package services

import "errors"

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrInvalidStatus     = errors.New("invalid status")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrConflict          = errors.New("conflict")
)
