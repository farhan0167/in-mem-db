package server

import (
	"farhan0167/mem-db/database"
	"log/slog"
	"net/http"
)

func NewServer(
	logger *slog.Logger,
	db *database.DB,
) http.Handler {

	mux := http.NewServeMux()

	addRoutes(mux, logger, db)
	return mux

}
