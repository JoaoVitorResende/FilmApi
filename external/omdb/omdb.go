package omdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"main/internal/application"
	"main/internal/exception"
)

type result struct {
	Search       []searchResult `json:"Search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"Response"`
	Error        string         `json:"Error"`
}

type searchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type Client struct {
	apiKey string
	http   *http.Client
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
		http:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (c Client) Search(ctx context.Context, name string) ([]application.FilmSummary, error) {
	v := url.Values{}
	v.Set("apikey", c.apiKey)
	v.Set("s", name)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://www.omdbapi.com/?"+v.Encode(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to build request: %v", exception.ErrUpstream, err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to call omdb: %v", exception.ErrUpstream, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: omdb returned status %d", exception.ErrUpstream, resp.StatusCode)
	}

	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("%w: failed to decode omdb response: %v", exception.ErrUpstream, err)
	}

	// OMDb responde 200 mesmo quando não encontra nada
	if res.Response == "False" {
		return nil, exception.ErrNotFound
	}

	films := make([]application.FilmSummary, 0, len(res.Search))
	for _, s := range res.Search {
		films = append(films, application.FilmSummary{
			Title:  s.Title,
			Year:   s.Year,
			ImdbID: s.ImdbID,
			Type:   s.Type,
			Poster: s.Poster,
		})
	}

	return films, nil
}