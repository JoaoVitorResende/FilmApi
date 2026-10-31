package domain

import (
	"main/internal/exception"
	"strings"
)

type Film struct {
	name string
}

func NewFilm(name string) (Film, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return Film{}, exception.ErrEmptyName
	}

	return Film{name: name}, nil
}

func (f Film) Name() string { return f.name }
