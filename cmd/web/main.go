package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	logger.Info("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		logger.Error("an unexpcted error occurred", "error", err)
		os.Exit(1)
	}
}
