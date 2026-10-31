package exception

import "errors"

var (
	ErrEmpty       = errors.New("item can't be empty")
	ErrEmptyName   = errors.New("film name can't be empty")
	ErrNotFound    = errors.New("item not found")
	ErrInvalidBody = errors.New("invalid body")
	ErrUpstream    = errors.New("communication failure")
)
