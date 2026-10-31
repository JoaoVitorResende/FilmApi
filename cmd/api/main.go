package main

import (
	"fmt"
	"log/slog"
	"main/external/omdb"
	"main/internal/application"
	"main/internal/interfaces/httpapi"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all system offline")
}

func run() error {
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		fmt.Println("variável de ambiente API_KEY não definida")
		os.Exit(1)
	}

	fmt.Println("chave carregada com sucesso")

	client := omdb.NewClient(apiKey)
	uc := application.NewSearchFilms(client)
	handler := httpapi.Handler(uc)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
