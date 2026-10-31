package application

import (
	"context"
	"errors"
	"testing"

	"main/internal/exception"
)

type fakeSearcher struct {
	films  []FilmSummary
	err    error
	called bool
}

func (f *fakeSearcher) Search(ctx context.Context, name string) ([]FilmSummary, error) {
	f.called = true
	return f.films, f.err
}

func TestSearchFilms_EmptyName(t *testing.T) {
	fake := &fakeSearcher{}
	uc := NewSearchFilms(fake)

	_, err := uc.Execute(context.Background(), "   ")

	if !errors.Is(err, exception.ErrEmptyName) {
		t.Fatalf("esperava ErrEmptyName, veio %v", err)
	}
	if fake.called {
		t.Fatal("não deveria chamar o OMDb com nome vazio")
	}
}

func TestSearchFilms_Success(t *testing.T) {
	fake := &fakeSearcher{films: []FilmSummary{{Title: "The Matrix"}}}
	uc := NewSearchFilms(fake)

	films, err := uc.Execute(context.Background(), "matrix")

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if len(films) != 1 || films[0].Title != "The Matrix" {
		t.Fatalf("resultado inesperado: %+v", films)
	}
}