package main

import (
	"farhan0167/mem-db/database"
	"farhan0167/mem-db/server"
	"log/slog"
	"net/http"
	"os"
)

func Run() error {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	).With("app", "api")

	db := database.DB{}

	http_hldr := server.NewServer(logger, &db)

	svr := &http.Server{
		Addr:    ":8080",
		Handler: http_hldr,
	}
	logger.Info("Starting server", "addr", svr.Addr)
	svr.ListenAndServe()

	return nil
}

func main() {
	err := Run()
	if err != nil {
		panic(err)
	}
}
