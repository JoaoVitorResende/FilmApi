package httpapi
//go test ./...
import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"main/internal/application"
	"main/internal/exception"
)

type stubSearcher struct {
	films []application.FilmSummary
	err   error
}

func (s stubSearcher) Search(ctx context.Context, name string) ([]application.FilmSummary, error) {
	return s.films, s.err
}

func TestGet_Status(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		stub   stubSearcher
		status int
	}{
		{"ok", "?s=matrix", stubSearcher{films: []application.FilmSummary{{Title: "The Matrix"}}}, http.StatusOK},
		{"nome vazio", "", stubSearcher{}, http.StatusBadRequest},
		{"não encontrado", "?s=xyz", stubSearcher{err: exception.ErrNotFound}, http.StatusNotFound},
		{"omdb fora", "?s=matrix", stubSearcher{err: exception.ErrUpstream}, http.StatusBadGateway},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := application.NewSearchFilms(tt.stub)
			h := Handler(uc)

			req := httptest.NewRequest(http.MethodGet, "/"+tt.query, nil)
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)

			if rec.Code != tt.status {
				t.Fatalf("esperava %d, veio %d", tt.status, rec.Code)
			}
		})
	}
}