package application

import (
	"context"
	"main/internal/domain"
)

type FilmSummary struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type FilmSearcher interface {
	Search(ctx context.Context, name string) ([]FilmSummary, error)
}

type SearchFilms struct {
	searcher FilmSearcher
}

func NewSearchFilms(s FilmSearcher) SearchFilms {
	return SearchFilms{searcher: s}
}

func (uc SearchFilms) Execute(ctx context.Context, name string) ([]FilmSummary, error) {
	film, err := domain.NewFilm(name)
	if err != nil {
		return nil, err
	}
	return uc.searcher.Search(ctx, film.Name())
}